package dto

type UpdateProjectNameRequest struct {
	Name string `json:"name" validate:"required"`
}
