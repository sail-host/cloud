package github

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"

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
