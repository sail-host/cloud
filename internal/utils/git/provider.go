package git

import (
	git "github.com/google/go-github/v65/github"
	"github.com/sail-host/cloud/internal/utils/git/github"
)

type GitProvider interface {
	CheckAccount() (bool, error)
	GetRepos(page, perPage int) (*github.ReposResponse, error)
	GetFramework(owner, repo string) (string, error)
	GetRepo(owner, repo string) (*git.Repository, error)
	GetLastCommitInBranch(owner, repo, branch string) (*git.RepositoryCommit, error)
	CloneRepo(owner, repo, path, branch, token, username string) error
	CreateDeployment(owner, repo string, deployment *git.DeploymentRequest) error
}
