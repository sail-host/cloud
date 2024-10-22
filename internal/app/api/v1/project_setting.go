package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sail-host/cloud/internal/app/dto"
)

func (b *BaseApi) UpdateProjectName(echo echo.Context) error {
	var request dto.UpdateProjectNameRequest
	var baseResponse dto.BaseResponse
	if err := echo.Bind(&request); err != nil {
		baseResponse.Status = "error"
		baseResponse.Message = err.Error()
		return echo.JSON(http.StatusBadRequest, baseResponse)
	}

	err := projectSettingService.UpdateProjectName(echo.Request().Context(), echo.Param("name"), request)
	if err != nil {
		baseResponse.Status = "error"
		baseResponse.Message = err.Error()

		return echo.JSON(http.StatusBadRequest, baseResponse)
	}

	baseResponse.Status = "success"
	baseResponse.Message = "Project name updated"

	return echo.JSON(http.StatusOK, baseResponse)
}

func (b *BaseApi) UpdateBuildAndOutputDir(echo echo.Context) error {
	var request dto.UpdateBuildAndOutputDirRequest
	var baseResponse dto.BaseResponse
	if err := echo.Bind(&request); err != nil {
		baseResponse.Status = "error"
		baseResponse.Message = err.Error()
		return echo.JSON(http.StatusBadRequest, baseResponse)
	}

	err := projectSettingService.UpdateBuildAndOutputDir(echo.Request().Context(), echo.Param("name"), request)
	if err != nil {
		baseResponse.Status = "error"
		baseResponse.Message = err.Error()
		return echo.JSON(http.StatusBadRequest, baseResponse)
	}

	baseResponse.Status = "success"
	baseResponse.Message = "Project build and output dir updated"

	return echo.JSON(http.StatusOK, baseResponse)
}

func (b *BaseApi) UpdateProjectDomain(echo echo.Context) error {
	var request dto.AddNewDomainRequest
	var baseResponse dto.BaseResponse
	if err := echo.Bind(&request); err != nil {
		baseResponse.Status = "error"
		baseResponse.Message = err.Error()
		return echo.JSON(http.StatusBadRequest, baseResponse)
	}

	err := echo.Validate(request)
	if err != nil {
		baseResponse.Status = "error"
		baseResponse.Message = err.Error()
		return echo.JSON(http.StatusBadRequest, baseResponse)
	}

	err = deploymentDomainService.AddNewDomain(echo.Param("name"), request)
	if err != nil {
		baseResponse.Status = "error"
		baseResponse.Message = err.Error()
		return echo.JSON(http.StatusBadRequest, baseResponse)
	}

	baseResponse.Status = "success"
	baseResponse.Message = "Project domain updated"

	return echo.JSON(http.StatusOK, baseResponse)
}
