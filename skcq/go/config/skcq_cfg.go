package config

import (
	"encoding/json"

	"go.skia.org/infra/go/skerr"
)

// SkCQCfg is a struct which describes the SkCQ config for a repo+branch at a
// particular commit.
type SkCQCfg struct {

	// Will use the internal skcq-fe instance to display results if true and
	// will use the skia.internal buildbucket. Cannot have both Internal and
	// Staging be true.
	Internal bool `json:"internal"`

	// Will use the staging skcq-fe instance to display results if true and
	// will use the skia.testing buildbucket. Cannot have both Internal and
	// Staging be true.
	Staging bool `json:"staging"`

	// Full path to tasks.json file to get list of CQ try jobs from.
	TasksJSONPath string `json:"tasks_json_path,omitempty"`

	// Name of the go/cria group that includes the list of people allowed to
	// commit to this repo+branch.
	CommitterList string `json:"committer_list"`

	// Name of the go/cria group that includes the list of people allowed to
	// run dry-runs on this repo+branch.
	DryRunAccessList string `json:"dry_run_access_list"`

	// The URL of the tree status instance that will gate submissions to this
	// repo+branch when the tree is closed.
	TreeStatusURL string `json:"tree_status_url,omitempty"`

	// The throttler config that will gate the rate of submissions to this
	// repo+branch.
	ThrottlerCfg *ThrottlerCfg `json:"throttler_cfg,omitempty"`
}

// ThrottlerCfg is a struct which describes how the rate of submissions to
// this repo+branch will be gated.
type ThrottlerCfg struct {
	// How many commits are allowed within BurstDelaySecs. Default used is
	// throttler.MaxBurstDefault.
	MaxBurst int `json:"max_burst"`
	// The window of seconds MaxBurst commits are allowed in. Default used is
	// throttler.BurstDelaySecs.
	BurstDelaySecs int `json:"burst_delay_secs"`
}

// Validate returns an error if the SkCQCfg is not valid.
func (c *SkCQCfg) Validate() error {
	if c.CommitterList == "" {
		return skerr.Fmt("Must specify a CommitterList")
	}
	if c.DryRunAccessList == "" {
		return skerr.Fmt("Must specify a DryRunAccessList")
	}
	if c.Internal && c.Staging {
		return skerr.Fmt("Cannot have both internal and staging be true")
	}
	return nil
}

// ParseSkCQCfg is a utility function that parses the given SkCQ cfg file
// contents and returns the config.
func ParseSkCQCfg(contents string) (*SkCQCfg, error) {
	var rv SkCQCfg
	if err := json.Unmarshal([]byte(contents), &rv); err != nil {
		return nil, skerr.Fmt("Failed to read SkCQ cfg: could not parse file: %s\nContents:\n%s", err, string(contents))
	}

	return &rv, nil
}