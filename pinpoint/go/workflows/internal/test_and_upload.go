package internal

import (
	"context"
	"time"

	buildbucketpb "go.chromium.org/luci/buildbucket/proto"
	"go.skia.org/infra/go/skerr"
	"go.skia.org/infra/go/sklog"
	"go.skia.org/infra/pinpoint/go/bot_configs"
	"go.skia.org/infra/pinpoint/go/clients/upload"
	"go.skia.org/infra/pinpoint/go/common"
	"go.skia.org/infra/pinpoint/go/workflows"
	pinpoint_proto "go.skia.org/infra/pinpoint/proto/v1"
	"go.temporal.io/sdk/workflow"
)

// TestResult represents the columns used by the BQ table.
type TestResult struct {
	// WorkflowID is the ID of the Temporal Workflow.
	WorkflowID string `bigquery:"workflow_id"`
	// Bot, also referred to as bot_configuration, is usually a bot name.
	// See bot_configs/internal.json or external.json for the fully supported list.
	Bot string `bigquery:"bot"`
	// SwarmingTaskID is the ID of the swarming task associated with the test run.
	SwarmingTaskID string `bigquery:"swarming_task_id"`
	// Benchmark is the benchmark name, ie/ Speedometer3.
	Benchmark string `bigquery:"benchmark"`
	// Story is the story associated with the benchmark. Not all benchmarks have
	// stories (ie/ Speedometer3)
	Story string `bigquery:"story"`
	// Chart is the sub module to the benchmark.
	Chart string `bigquery:"chart"`
	// SampleValues is the list of values generated by the benchmark run.
	SampleValues []float64 `bigquery:"sample_values"`
	// TaskFailed indicates whether the benchmark run failed.
	TaskFailed bool `bigquery:"task_failed"`
	// PGOEnabled indicates whether PGO was enabled for the Chrome build.
	PGOEnabled bool `bigquery:"pgo_enabled"`
	// CreateTime is when the row data is generated.
	CreateTime time.Time `bigquery:"create_time"`
}

// TestAndExportParams contains information needed to build a version of Chrome
// and trigger the tests on Swarming.
type TestAndExportParams struct {
	// WorkflowID is the ID of the Temporal Workflow.
	WorkflowID string
	// Benchmark is the benchmark name, ie/ speedometer3.
	Benchmark string
	// Bot, also referred to as bot_configuration, is usually a bot name.
	Bot string
	// Git Hash associated with chromium/src to build.
	GitHash string
	// Iterations is the number of runs to trigger.
	// Note: Temporal will struggle and crash with more than 10,000 steps. Each
	// activity generates 3, so be cautious when running more than 1,000 iterations.
	// If you need to run more than 1,000, you can run several of these Test
	// AndExportWorkflows.
	Iterations int
	// Project is the Google Cloud Project (ie/ chromeperf), used for creating the
	// fully qualified BQ table name.
	Project string
	// Dataset is the BQ Dataset. Format is {project}.{dataset}.
	// Used for creating the fully qualified BQ table name.
	Dataset string
	// TableName is the name of the table within Dataset. Format is
	// {project}.{dataset}.{tableName}. Used for creating the fully qualified BQ
	// table name.
	TableName string
}

// processInsertData takes the task IDs and serializes them into TestResult objects,
// which can be used to upload to BQ.
func processInsertData(req *TestAndExportParams, taskIds []string) []*TestResult {
	res := []*TestResult{}
	currentTime := time.Now().UTC()
	for _, taskId := range taskIds {
		res = append(res, &TestResult{
			WorkflowID:     req.WorkflowID,
			Benchmark:      req.Benchmark,
			Bot:            req.Bot,
			SwarmingTaskID: taskId,
			PGOEnabled:     true,
			CreateTime:     currentTime,
		})
	}

	return res
}

// build is a wrapper around BuildChrome workflow to build a version of chrome at the provided commit.
func buildArtifact(ctx workflow.Context, workflowID, bot, benchmark string, combinedCommit *common.CombinedCommit) (*workflows.Build, error) {
	t, err := bot_configs.GetIsolateTarget(bot, benchmark)
	if err != nil {
		return nil, skerr.Wrapf(err, "no target found for (%s, %s)", bot, benchmark)
	}

	var b *workflows.Build
	if err := workflow.ExecuteChildWorkflow(ctx, workflows.BuildChrome, workflows.BuildParams{
		WorkflowID: workflowID,
		Device:     bot,
		Target:     t,
		Commit:     combinedCommit,
		Project:    "chromium",
	}).Get(ctx, &b); err != nil {
		return nil, skerr.Wrap(err)
	}

	if b.Status != buildbucketpb.Status_SUCCESS {
		return nil, skerr.Fmt("build failed")
	}

	return b, nil
}

// test is a wrapper around ScheduleTaskActivity which runs the benchmark
// against the provided build artifact.
func test(ctx workflow.Context, p *RunBenchmarkParams) (string, error) {
	var rba RunBenchmarkActivity
	var taskID string
	if err := workflow.ExecuteActivity(ctx, rba.ScheduleTaskActivity, p).Get(ctx, &taskID); err != nil {
		return "", skerr.Wrap(err)
	}

	return taskID, nil
}

// uploadTestRuns will ensure that the target table exists, and
// upload all swarming task ids triggered for this workflow.
func uploadTestRuns(ctx workflow.Context, req *TestAndExportParams, taskIds []string) error {
	insertData := processInsertData(req, taskIds)

	var wferr error
	if err := workflow.ExecuteActivity(ctx, UploadResultsActivity, req.Project, req.Dataset, req.TableName, insertData).Get(ctx, &wferr); err != nil {
		return skerr.Wrapf(err, "failed activity upload")
	}
	return wferr
}

func UploadResultsActivity(ctx context.Context, project, dataset, tableName string, results []*TestResult) error {
	uploadClient, err := upload.NewUploadClient(ctx, &upload.UploadClientConfig{
		Project:   project,
		DatasetID: dataset,
		TableName: tableName,
	})

	if err != nil {
		return skerr.Wrapf(err, "Failed to create upload client")
	}

	if err := uploadClient.CreateTableFromStruct(ctx, &upload.CreateTableRequest{
		Definition: TestResult{},
	}); err != nil {
		sklog.Infof("Tables likely exist. Continuing with workflow")
	}

	if err := uploadClient.Insert(ctx, &upload.InsertRequest{
		Items: results,
	}); err != nil {
		return skerr.Wrapf(err, "Upload client failed to insert data.")
	}

	return nil
}

// RunTestAndExportWorkflow will build chrome, run tests and export the swarming
// task ids to BQ.
func RunTestAndExportWorkflow(ctx workflow.Context, req *TestAndExportParams) error {
	ctx = workflow.WithChildOptions(ctx, childWorkflowOptions)
	ctx = workflow.WithActivityOptions(ctx, regularActivityOptions)
	ctx = workflow.WithLocalActivityOptions(ctx, localActivityOptions)

	combinedCommit := common.NewCombinedCommit(&pinpoint_proto.Commit{GitHash: req.GitHash})

	// build chrome
	resp, err := buildArtifact(ctx, req.WorkflowID, req.Bot, req.Benchmark, combinedCommit)
	if err != nil {
		return skerr.Wrapf(err, "Failed to build Chrome")
	}

	testParams := &RunBenchmarkParams{
		JobID:     req.WorkflowID,
		BuildCAS:  resp.CAS,
		Commit:    combinedCommit,
		BotConfig: req.Bot,
		Benchmark: req.Benchmark,
	}

	// run tests x number of times, according to the defined iteration
	// count from the request.
	taskIds := []string{}
	failedCount := 0
	for i := 0; i < req.Iterations; i++ {
		taskId, err := test(ctx, testParams)
		if err != nil {
			failedCount++
		}

		taskIds = append(taskIds, taskId)
	}

	// upload the rows to BQ
	if err := uploadTestRuns(ctx, req, taskIds); err != nil {
		return skerr.Wrapf(err, "Failed to upload rows")
	}
	if failedCount > 0 {
		return skerr.Fmt("Failed to trigger %d benchmark runs", failedCount)
	}
	return nil
}