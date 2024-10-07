package github

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/google/go-github/v65/github"
)

type Github struct {
	Client *github.Client
	Owner  string
}

func NewGithub(token, owner string) *Github {
	client := github.NewClient(nil).WithAuthToken(token)
	return &Github{Client: client, Owner: owner}
}

// Check if the account is an organization
func (g *Github) IsOrganization() (bool, error) {
	ctx := context.Background()
	user, _, err := g.Client.Users.Get(ctx, g.Owner)
	if err != nil {
		return false, err
	}

	if user.GetType() == "Organization" {
		return true, nil
	}
	return false, nil
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

type ReposResponse struct {
	Repos    []*github.Repository
	LastPage int
	NextPage int
}

func (g *Github) GetRepos(page, perPage int) (*ReposResponse, error) {
	ctx := context.Background()

	isOrg, err := g.IsOrganization()
	if err != nil {
		return nil, err
	}

	var repos []*github.Repository
	var response *github.Response

	if isOrg {
		repos, response, err = g.Client.Repositories.ListByOrg(ctx, g.Owner, &github.RepositoryListByOrgOptions{
			ListOptions: github.ListOptions{
				Page:    page,
				PerPage: perPage,
			},
			Sort: "created",
		})
	} else {
		repos, response, err = g.Client.Repositories.ListByUser(ctx, g.Owner, &github.RepositoryListByUserOptions{
			ListOptions: github.ListOptions{
				Page:    page,
				PerPage: perPage,
			},
			Sort: "created",
		})
	}

	if err != nil {
		return nil, err
	}

	return &ReposResponse{
		Repos:    repos,
		LastPage: response.LastPage,
		NextPage: response.NextPage,
	}, nil
}

func (g *Github) GetFramework(owner, repo string) (string, error) {
	ctx := context.Background()
	content, _, _, err := g.Client.Repositories.GetContents(ctx, owner, repo, "package.json", nil)
	if err != nil {
		return "", err
	}

	if content == nil {
		return "", errors.New("content not found")
	}

	decodedContent, err := base64.StdEncoding.DecodeString(*content.Content)
	if err != nil {
		return "", err
	}

	contentStr := string(decodedContent)

	packageJson := make(map[string]interface{})
	err = json.Unmarshal([]byte(contentStr), &packageJson)
	if err != nil {
		return "", err
	}

	dependencies, ok := packageJson["dependencies"].(map[string]interface{})
	if !ok {
		return "", errors.New("dependencies not found or invalid format")
	}

	frameworks := []string{"next", "react", "nuxt", "vue", "svelte", "remix", "angular", "lit", "ember", "vanilla", "vite"}
	for _, framework := range frameworks {
		if _, exists := dependencies[framework]; exists {
			return framework, nil
		}
	}

	devDependencies, ok := packageJson["devDependencies"].(map[string]interface{})
	if ok {
		for _, framework := range frameworks {
			if _, exists := devDependencies[framework]; exists {
				return framework, nil
			}
		}
	}

	return "", errors.New("framework not found")
}

func (g *Github) GetRepo(owner, repo string) (*github.Repository, error) {
	ctx := context.Background()
	repository, _, err := g.Client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return nil, err
	}

	return repository, nil
}

func (g *Github) GetLastCommitInBranch(owner, repo, branch string) (*github.RepositoryCommit, error) {
	ctx := context.Background()

	commits, _, err := g.Client.Repositories.ListCommits(ctx, owner, repo, &github.CommitsListOptions{
		SHA: branch,
	})
	if err != nil {
		return nil, err
	}

	return commits[0], nil
}

func (g *Github) CloneRepo(owner, repo, path, branch, token, username string) error {
	ctx := context.Background()

	repository, _, err := g.Client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return err
	}

	cloneURL := repository.GetCloneURL()
	if cloneURL == "" {
		return errors.New("failed to get clone URL for repository")
	}

	_, err = git.PlainClone(path, false, &git.CloneOptions{
		URL:           cloneURL,
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
		SingleBranch:  true,
		Depth:         1,
		Auth: &http.BasicAuth{
			Username: username,
			Password: token,
		},
	})
	if err != nil {
		return fmt.Errorf("git clone failed: %w", err)
	}

	return nil
}
