package v1

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sail-host/cloud/internal/app/dto"
	"github.com/sail-host/cloud/internal/global"
)

func (b *BaseApi) GitInternalList(echo echo.Context) error {
	var baseError dto.BaseError
	var page int
	var err error
	id := echo.Param("id")
	page_str := echo.QueryParam("page")
	if page_str != "" {
		page, err = strconv.Atoi(page_str)
		if err != nil {
			baseError.Status = "error"
			baseError.Message = "Invalid page"
			return echo.JSON(http.StatusBadRequest, baseError)
		}
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		baseError.Status = "error"
		baseError.Message = "Invalid ID"
		return echo.JSON(http.StatusBadRequest, baseError)
	}

	repos, err := gitInternalService.GetRepos(uint(idUint), page)
	if err != nil {
		baseError.Status = "error"
		baseError.Message = "Failed to get repos"
		global.LOG.Error(err)
		return echo.JSON(http.StatusInternalServerError, baseError)
	}

	var response dto.GitInternalResponse
	response.Status = "success"
	response.Message = "Repos fetched successfully"
	response.Data = repos.Data
	response.NextPage = repos.NextPage
	response.LastPage = repos.LastPage

	return echo.JSON(http.StatusOK, response)
}
