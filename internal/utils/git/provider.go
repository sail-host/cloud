package git

import "github.com/google/go-github/v65/github"

type GitProvider interface {
	CheckAccount() (bool, error)
	GetRepos() ([]*github.Repository, error)
}
