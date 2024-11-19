package service

import (
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
)

type WebhookGithubService struct {
}

type IWebhookGithubService interface {
	HandleWebhook(c echo.Context) error
}

func NewIWebhookGithubService() IWebhookGithubService {
	return &WebhookGithubService{}
}

type GithubWebhookPayload struct {
	RefType string `json:"ref_type"`
	Ref     string `json:"ref"`
	Before  string `json:"before"`
	After   string `json:"after"`
	Created bool   `json:"created"`
	Deleted bool   `json:"deleted"`
	Pusher  struct {
		Name string `json:"name"`
	} `json:"pusher"`
	Repository struct {
		Name string `json:"name"`
	} `json:"repository"`
	Sender struct {
		Login string `json:"login"`
	} `json:"sender"`
}

func (s *WebhookGithubService) HandleWebhook(c echo.Context) error {
	var payload GithubWebhookPayload
	err := c.Bind(&payload)
	if err != nil {
		return err
	}

	// Get project id from param
	projectID := c.Param("project_id")
	if projectID == "" {
		return errors.New("project id is required")
	}

	id, err := strconv.ParseUint(projectID, 10, 64)
	if err != nil {
		return err
	}

	project, err := projectRepo.GetProjectByID(uint(id))
	if err != nil {
		return err
	}

	// Check project branch
	if payload.Ref != project.ProductionBranch {
		return errors.New("branch not match")
	}

	// Create new deployment
	deployService := NewIDeployService()
	go deployService.Deploy(project)

	return nil
}
