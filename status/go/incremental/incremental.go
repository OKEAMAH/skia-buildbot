package incremental

/*
   Allow incremental updates to the client.
*/

import (
	"context"
	"sort"
	"sync"
	"time"

	"go.skia.org/infra/go/git/gitinfo"
	"go.skia.org/infra/go/git/repograph"
	"go.skia.org/infra/go/metrics2"
	"go.skia.org/infra/go/sklog"
	"go.skia.org/infra/go/util"
	"go.skia.org/infra/go/vcsinfo"
	"go.skia.org/infra/task_scheduler/go/db"
	"go.skia.org/infra/task_scheduler/go/window"
)

// Task is a trimmed-down version of db.Task for minimizing the amount of data
// we send to the client.
type Task struct {
	Commits        []string      `json:"commits"`
	Name           string        `json:"name"`
	Id             string        `json:"id"`
	Revision       string        `json:"revision"`
	Status         db.TaskStatus `json:"status"`
	SwarmingTaskId string        `json:"swarming_task_id"`
}

// Update represents all of the new information we obtained in a single Update()
// tick. Every time Update() is called on IncrementalCache, a new Update object
// is stored internally. When the client calls any variant of Get, any new
// Updates are found and merged into a single Update object to return.
type Update struct {
	BranchHeads      []*gitinfo.GitBranch                 `json:"branch_heads,omitempty"`
	CommitComments   map[string][]*CommitComment          `json:"commit_comments,omitempty"`
	Commits          []*vcsinfo.LongCommit                `json:"commits,omitempty"`
	StartOver        *bool                                `json:"start_over,omitempty"`
	SwarmingUrl      string                               `json:"swarming_url,omitempty"`
	TaskComments     map[string]map[string][]*TaskComment `json:"task_comments,omitempty"`
	Tasks            []*Task                              `json:"tasks,omitempty"`
	TaskSchedulerUrl string                               `json:"task_scheduler_url,omitempty"`
	TaskSpecComments map[string][]*TaskSpecComment        `json:"task_spec_comments,omitempty"`
	Timestamp        time.Time                            `json:"-"`
}

// IncrementalCache is a cache used for sending only new information to a
// client. New data is obtained at each Update() tick and stored internally with
// a timestamp. When the client requests new data, we return a combined set of
// Updates.
type IncrementalCache struct {
	comments         *commentsCache
	commits          *commitsCache
	mtx              sync.RWMutex
	numCommits       int
	swarmingUrl      string
	taskSchedulerUrl string
	tasks            *taskCache
	// Updates, keyed by repo and sorted ascending by timestamp.
	updates map[string][]*Update
	w       *window.Window
}

// NewIncrementalCache returns an IncrementalCache instance.
func NewIncrementalCache(d db.RemoteDB, w *window.Window, repos repograph.Map, numCommits int, swarmingUrl, taskSchedulerUrl string) (*IncrementalCache, error) {
	c := &IncrementalCache{
		comments:         newCommentsCache(d, repos),
		commits:          newCommitsCache(repos),
		numCommits:       numCommits,
		swarmingUrl:      swarmingUrl,
		taskSchedulerUrl: taskSchedulerUrl,
		tasks:            newTaskCache(d),
		w:                w,
	}
	return c, c.Update()
}

// getUpdatesInRange is a helper function which retrieves all Update objects
// within a given time range.
func (c *IncrementalCache) getUpdatesInRange(repo string, from, to time.Time) []*Update {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	from = from.UTC()
	to = to.UTC()
	// Obtain all updates in the given range.
	updates := []*Update{}
	// TODO(borenet): Could use binary search to get to the starting point
	// faster.
	for _, u := range c.updates[repo] {
		if !u.Timestamp.Before(from) && u.Timestamp.Before(to) {
			updates = append(updates, u)
		}
	}
	return updates
}

// GetRange returns all newly-obtained data in the given time range, trimmed
// to maxCommits.
func (c *IncrementalCache) GetRange(repo string, from, to time.Time, maxCommits int) (*Update, error) {
	updates := c.getUpdatesInRange(repo, from, to)
	// Merge the updates.
	rv := &Update{
		BranchHeads: nil,
		Commits:     []*vcsinfo.LongCommit{},
		StartOver:   nil,
		Tasks:       []*Task{},
	}
	for _, u := range updates {
		if u.BranchHeads != nil {
			rv.BranchHeads = u.BranchHeads
		}
		if u.CommitComments != nil {
			rv.CommitComments = u.CommitComments
		}
		rv.Commits = append(rv.Commits, u.Commits...)
		if u.StartOver != nil && *u.StartOver {
			rv.StartOver = u.StartOver
		}
		if u.TaskComments != nil {
			rv.TaskComments = u.TaskComments
		}
		rv.Tasks = append(rv.Tasks, u.Tasks...)
		if u.TaskSpecComments != nil {
			rv.TaskSpecComments = u.TaskSpecComments
		}
	}
	// Limit to only the most recent N commits.
	sort.Sort(vcsinfo.LongCommitSlice(rv.Commits)) // Most recent first.
	if len(rv.Commits) > maxCommits {
		rv.Commits = rv.Commits[:maxCommits]
	}
	// Replace empty slices with nil to save a few bytes in transfer.
	if len(rv.Commits) == 0 {
		rv.Commits = nil
	}
	if len(rv.Tasks) == 0 {
		rv.Tasks = nil
	}
	// If rv.StartOver is true, then we're providing all of the data the
	// client needs and we can perform an additional fitering step on the
	// tasks to ensure that we don't send older tasks the client doesn't
	// care about.
	if rv.StartOver != nil && *rv.StartOver {
		commits := make(map[string]bool, len(rv.Commits))
		for _, c := range rv.Commits {
			commits[c.Hash] = true
		}
		filteredTasks := make([]*Task, 0, len(rv.Tasks))
		for _, t := range rv.Tasks {
			if commits[t.Revision] {
				filteredTasks = append(filteredTasks, t)
			}
		}
		rv.Tasks = filteredTasks

		// Also provide the Swarming and Task Scheduler URLs when
		// starting over.
		rv.SwarmingUrl = c.swarmingUrl
		rv.TaskSchedulerUrl = c.taskSchedulerUrl
	}
	return rv, nil
}

// Get returns all newly-obtained data since the given time, trimmed to
// maxComits.
func (c *IncrementalCache) Get(repo string, since time.Time, maxCommits int) (*Update, error) {
	return c.GetRange(repo, since, time.Now().UTC(), maxCommits)
}

// GetAll returns all of the data in the cache, trimmed to maxCommits.
func (c *IncrementalCache) GetAll(repo string, maxCommits int) (*Update, error) {
	return c.Get(repo, c.w.Start(repo), maxCommits)
}

// Update obtains new data and stores it internally keyed by the current time.
func (c *IncrementalCache) Update() error {
	now := time.Now().UTC()
	if err := c.w.Update(); err != nil {
		return err
	}
	comments, err := c.comments.Update(c.w)
	if err != nil {
		return err
	}
	// TODO(borenet): If anything below here fails, the new tasks will be
	// lost forever!
	newTasks, startOver, err := c.tasks.Update(c.w)
	if err != nil {
		return err
	}
	branchHeads, commits, err := c.commits.Update(c.w, startOver, c.numCommits)
	if err != nil {
		return err
	}
	if startOver {
		c.comments.Reset()
		comments, err = c.comments.Update(c.w)
		if err != nil {
			return err
		}
	}
	updates := map[string]*Update{}
	var so *bool
	if startOver {
		so = new(bool)
		*so = true
	}
	for _, repo := range c.comments.repos {
		tasks := newTasks[repo]
		rc := comments[repo]
		haveNewData := len(branchHeads[repo]) > 0 ||
			len(commits[repo]) > 0 ||
			len(tasks) > 0 ||
			len(rc.CommitComments) > 0 ||
			len(rc.TaskComments) > 0 ||
			len(rc.TaskSpecComments) > 0
		if haveNewData {
			updates[repo] = &Update{
				BranchHeads:      branchHeads[repo],
				CommitComments:   rc.CommitComments,
				Commits:          commits[repo],
				StartOver:        so,
				TaskComments:     rc.TaskComments,
				Tasks:            tasks,
				TaskSpecComments: rc.TaskSpecComments,
				Timestamp:        now,
			}
		}
	}
	c.mtx.Lock()
	defer c.mtx.Unlock()
	if startOver {
		c.updates = map[string][]*Update{}
	}
	for repo, u := range updates {
		c.updates[repo] = append(c.updates[repo], u)
	}
	return nil
}

// UpdateLoop runs c.Update() in a loop.
func (c *IncrementalCache) UpdateLoop(frequency time.Duration, ctx context.Context) {
	lv := metrics2.NewLiveness("last_successful_incremental_cache_update")
	go util.RepeatCtx(frequency, ctx, func() {
		if err := c.Update(); err != nil {
			sklog.Errorf("Failed to update incremental cache: %s", err)
		} else {
			lv.Reset()
		}
	})
}
