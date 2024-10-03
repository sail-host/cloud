package git

import "github.com/google/go-github/v65/github"

type GitManager struct {
	Provider GitProvider
}

func (gm *GitManager) CheckAccount() (bool, error) {
	return gm.Provider.CheckAccount()
}

func (gm *GitManager) GetRepos() ([]*github.Repository, error) {
	return gm.Provider.GetRepos()
}

func (gm *GitManager) GetFramework(owner, repo string) (string, error) {
	return gm.Provider.GetFramework(owner, repo)
}

func NewGitManager(provider GitProvider) *GitManager {
	return &GitManager{Provider: provider}
}
