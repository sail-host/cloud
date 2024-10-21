package git

import (
	git "github.com/google/go-github/v65/github"
	"github.com/sail-host/cloud/internal/utils/git/github"
)

type GitManager struct {
	Provider GitProvider
}

func (gm *GitManager) CheckAccount() (bool, error) {
	return gm.Provider.CheckAccount()
}

func (gm *GitManager) GetRepos(page, perPage int) (*github.ReposResponse, error) {
	return gm.Provider.GetRepos(page, perPage)
}

func (gm *GitManager) GetFramework(owner, repo string) (string, error) {
	return gm.Provider.GetFramework(owner, repo)
}

func (gm *GitManager) GetRepo(owner, repo string) (*git.Repository, error) {
	return gm.Provider.GetRepo(owner, repo)
}

func (gm *GitManager) GetLastCommitInBranch(owner, repo, branch string) (*git.RepositoryCommit, error) {
	return gm.Provider.GetLastCommitInBranch(owner, repo, branch)
}

func (gm *GitManager) CloneRepo(owner, repo, path, branch, token, username string) error {
	return gm.Provider.CloneRepo(owner, repo, path, branch, token, username)
}

func (gm *GitManager) CreateDeployment(owner, repo string, deployment *git.DeploymentRequest) error {
	return gm.Provider.CreateDeployment(owner, repo, deployment)
}

func (gm *GitManager) UpdateDeploymentStatus(owner, repo, status, message string, deploymentID int64) error {
	return gm.Provider.UpdateDeploymentStatus(owner, repo, status, message, deploymentID)
}

func NewGitManager(provider GitProvider) *GitManager {
	return &GitManager{Provider: provider}
}
