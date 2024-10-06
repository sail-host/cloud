package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sail-host/cloud/internal/app/dto"
)

func (b *BaseApi) CreateProject(echo echo.Context) error {
	var request dto.CreateProjectRequest
	var baseError dto.BaseError
	if err := echo.Bind(&request); err != nil {
		baseError.Status = "error"
		baseError.Message = err.Error()
		return echo.JSON(http.StatusBadRequest, baseError)
	}
	err := projectService.CreateProject(echo, &request)
	if err != nil {
		baseError.Status = "error"
		baseError.Message = err.Error()
		return echo.JSON(http.StatusBadRequest, baseError)
	}
	var baseResponse dto.BaseResponse
	baseResponse.Status = "success"
	baseResponse.Message = "Project creation started"
	return echo.JSON(http.StatusOK, baseResponse)
}

func (b *BaseApi) ListProjects(echo echo.Context) error {
	projects, err := projectService.ListProjects()
	if err != nil {
		var baseError dto.BaseError
		baseError.Status = "error"
		baseError.Message = err.Error()
		return echo.JSON(http.StatusBadRequest, baseError)
	}
	return echo.JSON(http.StatusOK, projects)
}
