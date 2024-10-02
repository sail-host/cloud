package github

import (
	"context"

	"github.com/google/go-github/v65/github"
)

type Github struct {
	Client *github.Client
}

func NewGithub(token string) *Github {
	client := github.NewClient(nil).WithAuthToken(token)
	return &Github{Client: client}
}

func (g *Github) CheckAccount() (bool, error) {
	user, _, err := g.Client.Users.Get(context.Background(), "")
	if err != nil {
		return false, err
	}

	if user.Login == nil {
		return false, nil
	}

	return true, nil
}

func (g *Github) GetRepos() ([]*github.Repository, error) {
	ctx := context.Background()
	repos, _, err := g.Client.Repositories.ListByAuthenticatedUser(ctx, &github.RepositoryListByAuthenticatedUserOptions{})
	if err != nil {
		return nil, err
	}

	return repos, nil
}
