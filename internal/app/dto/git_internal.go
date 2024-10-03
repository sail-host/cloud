package dto

import "time"

type GitInternalResponse struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    []GitInternalRepo `json:"data"`
}

type GitInternalRepo struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	FullName      string    `json:"full_name"`
	Description   string    `json:"description"`
	URL           string    `json:"url"`
	DefaultBranch string    `json:"default_branch"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	CloneURL      string    `json:"clone_url"`
	Private       bool      `json:"private"`
	Framework     string    `json:"framework"`
}
