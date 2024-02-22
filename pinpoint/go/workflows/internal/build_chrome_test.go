package internal

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"go.skia.org/infra/pinpoint/go/workflows"

	"github.com/stretchr/testify/require"
	buildbucketpb "go.chromium.org/luci/buildbucket/proto"
	swarmingV1 "go.chromium.org/luci/common/api/swarming/swarming/v1"
	"go.temporal.io/sdk/testsuite"
)

func Test_BuildChrome_ShouldReturnBuild(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	var bca *BuildChromeActivity
	const fakeBuildID = int64(1234)
	cas := &swarmingV1.SwarmingRpcsCASReference{
		CasInstance: "fake-instance",
	}

	env.OnActivity(bca.SearchOrBuildActivity, mock.Anything, mock.Anything).Return(fakeBuildID, nil).Once()
	env.OnActivity(bca.WaitBuildCompletionActivity, mock.Anything, fakeBuildID).Return(buildbucketpb.Status_SUCCESS, nil).Once()
	env.OnActivity(bca.RetrieveCASActivity, mock.Anything, mock.Anything, mock.Anything).Return(cas, nil).Once()

	env.ExecuteWorkflow(BuildChrome, workflows.BuildChromeParams{})

	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
	var result *workflows.Build
	require.NoError(t, env.GetWorkflowResult(&result))
	require.Equal(t, &workflows.Build{
		ID:     fakeBuildID,
		Status: buildbucketpb.Status_SUCCESS,
		CAS:    cas,
	}, result)
	env.AssertExpectations(t)
}

func Test_BuildChrome_ShouldPopulateBuildError(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	var bca *BuildChromeActivity
	const fakeBuildID = int64(1234)

	env.OnActivity(bca.SearchOrBuildActivity, mock.Anything, mock.Anything).Return(fakeBuildID, nil)
	env.OnActivity(bca.WaitBuildCompletionActivity, mock.Anything, fakeBuildID).Return(buildbucketpb.Status_FAILURE, nil).Once()
	env.OnActivity(bca.RetrieveCASActivity, mock.Anything, mock.Anything, mock.Anything).Never()

	env.ExecuteWorkflow(BuildChrome, workflows.BuildChromeParams{})

	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
	var result *workflows.Build
	require.NoError(t, env.GetWorkflowResult(&result))
	require.Equal(t, &workflows.Build{
		ID:     fakeBuildID,
		Status: buildbucketpb.Status_FAILURE,
		CAS:    nil,
	}, result)
	env.AssertExpectations(t)
}
