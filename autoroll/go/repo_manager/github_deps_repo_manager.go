package repo_manager

import (
	"context"
	"net/http"
	"path/filepath"

	"go.skia.org/infra/autoroll/go/codereview"
	"go.skia.org/infra/autoroll/go/config_vars"
	"go.skia.org/infra/autoroll/go/repo_manager/child"
	"go.skia.org/infra/autoroll/go/repo_manager/common/git_common"
	"go.skia.org/infra/autoroll/go/repo_manager/common/version_file_common"
	"go.skia.org/infra/autoroll/go/repo_manager/parent"
	"go.skia.org/infra/go/depot_tools/deps_parser"
	"go.skia.org/infra/go/git"
	"go.skia.org/infra/go/github"
	"go.skia.org/infra/go/skerr"
)

// GithubDEPSRepoManagerConfig provides configuration for the Github RepoManager.
type GithubDEPSRepoManagerConfig struct {
	DepotToolsRepoManagerConfig
	ChildRepo   string `json:"childRepo"`
	ForkRepoURL string `json:"forkRepoURL"`
	// Optional config to use if parent path is different than
	// workdir + parent repo.
	GithubParentPath string `json:"githubParentPath,omitempty"`

	// Optional; transitive dependencies to roll. This is a mapping of
	// dependencies of the child repo which are also dependencies of the
	// parent repo and should be rolled at the same time. Keys are paths
	// to transitive dependencies within the child repo (as specified in
	// DEPS), and values are paths to those dependencies within the parent
	// repo.
	TransitiveDeps map[string]string `json:"transitiveDeps"`
}

// Validate the config.
func (c *GithubDEPSRepoManagerConfig) Validate() error {
	_, _, err := c.splitParentChild()
	if err != nil {
		return skerr.Wrap(err)
	}
	if c.ForkRepoURL == "" {
		return skerr.Fmt("ForkRepoURL is required")
	}
	return nil
}

// splitParentChild splits the GithubDEPSRepoManagerConfig into a
// parent.DEPSLocalConfig and a child.GitCheckoutConfig.
// TODO(borenet): Update the config format to directly define the parent
// and child. We shouldn't need most of the New.*RepoManager functions.
func (c GithubDEPSRepoManagerConfig) splitParentChild() (parent.DEPSLocalConfig, child.GitCheckoutConfig, error) {
	parentCfg := parent.DEPSLocalConfig{
		GitCheckoutConfig: parent.GitCheckoutConfig{
			BaseConfig: parent.BaseConfig{
				ChildPath:       c.DepotToolsRepoManagerConfig.CommonRepoManagerConfig.ChildPath,
				ChildRepo:       c.ChildRepo,
				IncludeBugs:     c.DepotToolsRepoManagerConfig.CommonRepoManagerConfig.IncludeBugs,
				IncludeLog:      c.DepotToolsRepoManagerConfig.CommonRepoManagerConfig.IncludeLog,
				CommitMsgTmpl:   c.DepotToolsRepoManagerConfig.CommonRepoManagerConfig.CommitMsgTmpl,
				MonorailProject: c.DepotToolsRepoManagerConfig.CommonRepoManagerConfig.BugProject,
			},
			GitCheckoutConfig: git_common.GitCheckoutConfig{
				Branch:  c.DepotToolsRepoManagerConfig.CommonRepoManagerConfig.ParentBranch,
				RepoURL: c.DepotToolsRepoManagerConfig.CommonRepoManagerConfig.ParentRepo,
			},
		},
		DependencyConfig: version_file_common.DependencyConfig{
			VersionFileConfig: version_file_common.VersionFileConfig{
				ID:   c.ChildRepo,
				Path: deps_parser.DepsFileName,
			},
		},
		CheckoutPath:   c.GithubParentPath,
		GClientSpec:    c.GClientSpec,
		PreUploadSteps: c.DepotToolsRepoManagerConfig.CommonRepoManagerConfig.PreUploadSteps,
		RunHooks:       c.RunHooks,
	}
	if err := parentCfg.Validate(); err != nil {
		return parent.DEPSLocalConfig{}, child.GitCheckoutConfig{}, skerr.Wrapf(err, "generated parent config is invalid")
	}
	childCfg := child.GitCheckoutConfig{
		GitCheckoutConfig: git_common.GitCheckoutConfig{
			Branch:      c.ChildBranch,
			RepoURL:     c.ChildRepo,
			RevLinkTmpl: c.DepotToolsRepoManagerConfig.CommonRepoManagerConfig.ChildRevLinkTmpl,
		},
	}
	if err := childCfg.Validate(); err != nil {
		return parent.DEPSLocalConfig{}, child.GitCheckoutConfig{}, skerr.Wrapf(err, "generated child config is invalid")
	}
	return parentCfg, childCfg, nil
}

// NewGithubDEPSRepoManager returns a RepoManager instance which operates in the given
// working directory and updates at the given frequency.
func NewGithubDEPSRepoManager(ctx context.Context, c *GithubDEPSRepoManagerConfig, reg *config_vars.Registry, workdir, rollerName string, githubClient *github.GitHub, recipeCfgFile, serverURL string, client *http.Client, cr codereview.CodeReview, local bool) (*parentChildRepoManager, error) {
	if err := c.Validate(); err != nil {
		return nil, skerr.Wrap(err)
	}
	parentCfg, childCfg, err := c.splitParentChild()
	if err != nil {
		return nil, skerr.Wrap(err)
	}
	uploadRoll := parent.GitCheckoutUploadGithubRollFunc(githubClient, cr.UserName(), rollerName)
	parentRM, err := parent.NewDEPSLocal(ctx, parentCfg, reg, client, serverURL, workdir, cr.UserName(), cr.UserEmail(), recipeCfgFile, uploadRoll)
	if err != nil {
		return nil, skerr.Wrap(err)
	}
	if err := parent.SetupGithub(ctx, parentRM, c.ForkRepoURL); err != nil {
		return nil, skerr.Wrap(err)
	}

	// Find the path to the child repo.
	childPath := filepath.Join(workdir, parentCfg.ChildPath)
	childCheckout := &git.Checkout{GitDir: git.GitDir(childPath)}
	childRM, err := child.NewGitCheckout(ctx, childCfg, reg, workdir, cr.UserName(), cr.UserEmail(), childCheckout)
	if err != nil {
		return nil, skerr.Wrap(err)
	}
	return newParentChildRepoManager(ctx, parentRM, childRM)
}
