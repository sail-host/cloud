package git

type GitManager struct {
	Provider GitProvider
}

func (gm *GitManager) CheckAccount() (bool, error) {
	return gm.Provider.CheckAccount()
}

func (gm *GitManager) GetRepos() ([]Repository, error) {
	return gm.Provider.GetRepos()
}

func NewGitManager(provider GitProvider) *GitManager {
	return &GitManager{Provider: provider}
}
