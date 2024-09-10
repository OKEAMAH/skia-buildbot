package analyzer

import (
	"sort"
	"sync"

	cpb "go.skia.org/infra/cabe/go/proto"

	apipb "go.chromium.org/luci/swarming/proto/api_v2"
)

// SwarmingTaskDiagnostics contains task-specific diagnostic messages generated by the Analyzer.
type SwarmingTaskDiagnostics struct {
	TaskID  string
	Message []string `json:",omitempty"`
}

// ReplicaDiagnostics contains replica, or task pair-specific diagnostic messages generated by the Analyzer.
type ReplicaDiagnostics struct {
	Number          int
	ControlTaskID   string
	TreatmentTaskID string
	Message         []string `json:",omitempty"`
}

// Diagnostics contains diagnostic messages about the replica task pairs and individual tasks generated
// by the Analyzer.
type Diagnostics struct {
	mu *sync.Mutex
	// Bad news: things that had to be excluded from the analysis, and why.
	ExcludedSwarmingTasks map[string]*SwarmingTaskDiagnostics `json:",omitempty"`
	ExcludedReplicas      map[int]*ReplicaDiagnostics         `json:",omitempty"`

	// Good news: things that were included in the analysis.
	IncludedSwarmingTasks map[string]*SwarmingTaskDiagnostics `json:",omitempty"`
	IncludedReplicas      map[int]*ReplicaDiagnostics         `json:",omitempty"`
}

func newDiagnostics() *Diagnostics {
	return &Diagnostics{
		mu:                    &sync.Mutex{},
		ExcludedSwarmingTasks: map[string]*SwarmingTaskDiagnostics{},
		ExcludedReplicas:      map[int]*ReplicaDiagnostics{},
		IncludedSwarmingTasks: map[string]*SwarmingTaskDiagnostics{},
		IncludedReplicas:      map[int]*ReplicaDiagnostics{},
	}
}

func (d *Diagnostics) excludeSwarmingTask(taskInfo *apipb.TaskRequestMetadataResponse, msg string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	taskDiag := d.ExcludedSwarmingTasks[taskInfo.TaskId]
	if taskDiag == nil {
		taskDiag = &SwarmingTaskDiagnostics{
			TaskID:  taskInfo.TaskId,
			Message: []string{},
		}
		d.ExcludedSwarmingTasks[taskInfo.TaskId] = taskDiag
	}
	taskDiag.Message = append(taskDiag.Message, msg)
}

func (d *Diagnostics) includeSwarmingTask(taskInfo *apipb.TaskRequestMetadataResponse) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.IncludedSwarmingTasks[taskInfo.TaskId] = &SwarmingTaskDiagnostics{
		TaskID: taskInfo.TaskId,
	}
}

func (d *Diagnostics) excludeReplica(replicaNumber int, pair pairedTasks, msg string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	replicaDiag := d.ExcludedReplicas[replicaNumber]
	if replicaDiag == nil {
		replicaDiag = &ReplicaDiagnostics{
			Number:          replicaNumber,
			ControlTaskID:   pair.control.taskID,
			TreatmentTaskID: pair.treatment.taskID,
			Message:         []string{},
		}
		d.ExcludedReplicas[replicaNumber] = replicaDiag
	}

	replicaDiag.Message = append(replicaDiag.Message, msg)
}

func (d *Diagnostics) includeReplica(replicaNumber int, pair pairedTasks) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.IncludedReplicas[replicaNumber] = &ReplicaDiagnostics{
		Number:          replicaNumber,
		ControlTaskID:   pair.control.taskID,
		TreatmentTaskID: pair.treatment.taskID,
	}
}

type bySwarmingTaskId []*cpb.SwarmingTaskDiagnostics

func (a bySwarmingTaskId) Len() int      { return len(a) }
func (a bySwarmingTaskId) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a bySwarmingTaskId) Less(i, j int) bool {
	if a[i].Id.Project != a[j].Id.Project {
		return a[i].Id.Project < a[j].Id.Project
	}
	return a[i].Id.TaskId < a[j].Id.TaskId
}

type byReplicaNumber []*cpb.ReplicaDiagnostics

func (a byReplicaNumber) Len() int      { return len(a) }
func (a byReplicaNumber) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byReplicaNumber) Less(i, j int) bool {
	return a[i].ReplicaNumber < a[j].ReplicaNumber
}

func (d *Diagnostics) AnalysisDiagnostics() *cpb.AnalysisDiagnostics {
	ret := &cpb.AnalysisDiagnostics{}
	for taskId, taskDiag := range d.ExcludedSwarmingTasks {
		ret.ExcludedSwarmingTasks = append(ret.ExcludedSwarmingTasks, &cpb.SwarmingTaskDiagnostics{
			Id: &cpb.SwarmingTaskId{
				TaskId: taskId,
			},
			Message: taskDiag.Message,
		})
	}
	for replicaNumber, replicaDiag := range d.ExcludedReplicas {
		ret.ExcludedReplicas = append(ret.ExcludedReplicas, &cpb.ReplicaDiagnostics{
			ReplicaNumber: int32(replicaNumber),
			ControlTask: &cpb.SwarmingTaskId{
				TaskId: replicaDiag.ControlTaskID,
			},
			TreatmentTask: &cpb.SwarmingTaskId{
				TaskId: replicaDiag.TreatmentTaskID,
			},
			Message: replicaDiag.Message,
		})
	}
	for taskId, taskDiag := range d.IncludedSwarmingTasks {
		ret.IncludedSwarmingTasks = append(ret.IncludedSwarmingTasks, &cpb.SwarmingTaskDiagnostics{
			Id: &cpb.SwarmingTaskId{
				TaskId: taskId,
			},
			Message: taskDiag.Message,
		})
	}
	for replicaNumber, replicaDiag := range d.IncludedReplicas {
		ret.IncludedReplicas = append(ret.IncludedReplicas, &cpb.ReplicaDiagnostics{
			ReplicaNumber: int32(replicaNumber),
			ControlTask: &cpb.SwarmingTaskId{
				TaskId: replicaDiag.ControlTaskID,
			},
			TreatmentTask: &cpb.SwarmingTaskId{
				TaskId: replicaDiag.TreatmentTaskID,
			},
			Message: replicaDiag.Message,
		})
	}
	sort.Sort(bySwarmingTaskId(ret.ExcludedSwarmingTasks))
	sort.Sort(byReplicaNumber(ret.ExcludedReplicas))
	sort.Sort(bySwarmingTaskId(ret.IncludedSwarmingTasks))
	sort.Sort(byReplicaNumber(ret.IncludedReplicas))
	return ret
}
