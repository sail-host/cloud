package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sail-host/cloud/internal/app/dto"
)

func (b *BaseApi) CreateProject(echo echo.Context) error {
	var request dto.CreateProjectRequest
	if err := echo.Bind(&request); err != nil {
		return echo.JSON(http.StatusBadRequest, err)
	}
	err := projectService.CreateProject(echo, &request)
	if err != nil {
		return echo.JSON(http.StatusBadRequest, err)
	}
	var baseResponse dto.BaseResponse
	baseResponse.Status = "success"
	baseResponse.Message = "Project creation started"
	return echo.JSON(http.StatusOK, baseResponse)
}
