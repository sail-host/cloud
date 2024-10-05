package git

import "github.com/sail-host/cloud/internal/utils/git/github"

type GitProvider interface {
	CheckAccount() (bool, error)
	GetRepos(page, perPage int) (*github.ReposResponse, error)
	GetFramework(owner, repo string) (string, error)
}
