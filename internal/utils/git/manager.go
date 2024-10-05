package git

import "github.com/sail-host/cloud/internal/utils/git/github"

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

func NewGitManager(provider GitProvider) *GitManager {
	return &GitManager{Provider: provider}
}
