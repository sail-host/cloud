package github

import (
	"context"

	"github.com/google/go-github/v65/github"
	"github.com/sail-host/cloud/internal/utils/git"
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

func (g *Github) GetRepos() ([]git.Repository, error) {
	var res []git.Repository
	ctx := context.Background()
	repos, _, err := g.Client.Repositories.ListAll(ctx, &github.RepositoryListAllOptions{})
	if err != nil {
		return nil, err
	}

	for _, r := range repos {
		res = append(res, git.Repository{
			ID:   *r.ID,
			Name: *r.Name,
			Url:  *r.URL,
		})
	}

	return res, nil
}
