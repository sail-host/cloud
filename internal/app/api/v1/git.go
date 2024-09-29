package v1

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sail-host/cloud/internal/app/dto"
)

func (b *BaseApi) GitList(echo echo.Context) error {

	gitList, err := gitService.GetListGit()
	if err != nil {
		return echo.JSON(http.StatusBadRequest, err)
	}
	return echo.JSON(http.StatusOK, gitList)
}

func (b *BaseApi) GitCreate(echo echo.Context) error {
	var request dto.CreateGitRequest
	if err := echo.Bind(&request); err != nil {
		return echo.JSON(http.StatusBadRequest, err)
	}
	git, err := gitService.CreateGit(echo, request)
	if err != nil {
		return echo.JSON(http.StatusBadRequest, err)
	}
	return echo.JSON(http.StatusOK, git)
}

func (b *BaseApi) GitUpdate(echo echo.Context) error {
	var baseError dto.BaseError
	id, err := strconv.ParseUint(echo.Param("id"), 10, 64)
	if err != nil {
		baseError.Status = "error"
		baseError.Message = "Invalid ID"
		return echo.JSON(http.StatusBadRequest, baseError)
	}
	var request dto.UpdateGitRequest
	if err := echo.Bind(&request); err != nil {
		baseError.Status = "error"
		baseError.Message = "Invalid request"
		return echo.JSON(http.StatusBadRequest, baseError)
	}
	git, baseErr := gitService.UpdateGit(echo, uint(id), request)
	if baseErr != nil {
		baseError.Status = "error"
		baseError.Message = baseErr.Message
		return echo.JSON(http.StatusBadRequest, baseError)
	}
	return echo.JSON(http.StatusOK, git)
}

func (b *BaseApi) GitDelete(echo echo.Context) error {
	var baseError dto.BaseError

	id, err := strconv.ParseUint(echo.Param("id"), 10, 64)
	if err != nil {
		baseError.Status = "error"
		baseError.Message = "Invalid ID"
		return echo.JSON(http.StatusBadRequest, baseError)
	}
	baseErr := gitService.DeleteGit(uint(id))
	if baseErr != nil {
		baseError.Status = "error"
		baseError.Message = baseErr.Message
		return echo.JSON(http.StatusBadRequest, baseError)
	}

	return echo.JSON(http.StatusNoContent, nil)
}

func (b *BaseApi) GitGetByID(echo echo.Context) error {
	var baseError dto.BaseError
	id, err := strconv.ParseUint(echo.Param("id"), 10, 64)
	if err != nil {
		baseError.Status = "error"
		baseError.Message = "Invalid ID"
		return echo.JSON(http.StatusBadRequest, baseError)
	}
	git, baseErr := gitService.GetGitByID(uint(id))
	if baseErr != nil {
		baseError.Status = "error"
		baseError.Message = baseErr.Message
		return echo.JSON(http.StatusBadRequest, baseError)
	}
	return echo.JSON(http.StatusOK, git)
}

func (b *BaseApi) GitCheckAccount(echo echo.Context) error {
	var request dto.CreateGitRequest
	if err := echo.Bind(&request); err != nil {
		return echo.JSON(http.StatusBadRequest, err)
	}

	git, err := gitService.CheckAccount(echo, request)
	if err != nil {
		return echo.JSON(http.StatusBadRequest, err)
	}
	return echo.JSON(http.StatusOK, git)
}
