package dto

import "github.com/sail-host/cloud/internal/app/model"

type GitResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Url   string `json:"url"`
	Type  string `json:"type"`
	Token string `json:"token"`
}

type GitListResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    []model.Git `json:"data"`
}

type CreateGitRequest struct {
	Name  string `json:"name" validate:"required"`
	Url   string `json:"url" validate:"required"`
	Type  string `json:"type" validate:"required"`
	Token string `json:"token" validate:"required"`
}

type UpdateGitRequest struct {
	Name  string `json:"name" validate:"required"`
	Url   string `json:"url" validate:"required"`
	Type  string `json:"type" validate:"required"`
	Token string `json:"token" validate:"required"`
}

type GetGitByIDRequest struct {
	ID uint `json:"id"`
}
