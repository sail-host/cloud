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
	err := deployService.CreateProject(echo, &request)
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

func (b *BaseApi) GetProjectWithName(echo echo.Context) error {
	projectName := echo.Param("name")
	project, err := projectService.GetProjectWithName(projectName)
	if err != nil {
		var baseError dto.BaseError
		baseError.Status = "error"
		baseError.Message = err.Error()
		return echo.JSON(http.StatusBadRequest, baseError)
	}

	return echo.JSON(http.StatusOK, project)
}

func (b *BaseApi) CheckProjectName(echo echo.Context) error {
	projectName := echo.QueryParam("name")
	project, err := projectService.CheckProjectName(projectName)
	if err != nil {
		var baseError dto.BaseError
		baseError.Status = "error"
		baseError.Message = err.Error()
		return echo.JSON(http.StatusBadRequest, baseError)
	}

	return echo.JSON(http.StatusOK, project)
}

func (b *BaseApi) GetProjectDeployments(echo echo.Context) error {
	projectName := echo.Param("name")
	deployments, err := projectService.GetProjectDeployments(projectName)
	if err != nil {
		var baseError dto.BaseError
		baseError.Status = "error"
		baseError.Message = err.Error()
		return echo.JSON(http.StatusBadRequest, baseError)
	}

	var baseResponse dto.BaseResponse
	baseResponse.Status = "success"
	baseResponse.Message = "Deployments fetched"
	baseResponse.Data = deployments

	return echo.JSON(http.StatusOK, baseResponse)
}

func (b *BaseApi) RedeployProject(echo echo.Context) error {
	projectName := echo.Param("name")
	err := deployService.Redeploy(echo, projectName)
	if err != nil {
		var baseError dto.BaseError
		baseError.Status = "error"
		baseError.Message = err.Error()
		return echo.JSON(http.StatusBadRequest, baseError)
	}

	var baseResponse dto.BaseResponse
	baseResponse.Status = "success"
	baseResponse.Message = "Project redeploy started"

	return echo.JSON(http.StatusOK, baseResponse)
}
