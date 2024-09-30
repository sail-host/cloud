package v1

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sail-host/cloud/internal/app/dto"
)

func (api *BaseApi) CreateDomain(c echo.Context) error {
	var baseError dto.BaseError
	request := dto.CreateDomainRequest{}

	if err := c.Bind(&request); err != nil {
		baseError.Status = "error"
		baseError.Message = err.Error()
		return c.JSON(http.StatusBadRequest, baseError)
	}

	res, err := domainService.CreateDomain(c, request)
	if err != nil && err.Status == "error" {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, res)
}

func (api *BaseApi) UpdateDomain(c echo.Context) error {
	var baseError dto.BaseError

	request := dto.UpdateDomainRequest{}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		baseError.Status = "error"
		baseError.Message = err.Error()
		return c.JSON(http.StatusBadRequest, baseError)
	}

	if err := c.Bind(&request); err != nil {
		baseError.Status = "error"
		baseError.Message = err.Error()
		return c.JSON(http.StatusBadRequest, baseError)
	}

	res, errBase := domainService.UpdateDomain(c, uint(id), request)
	if errBase != nil && errBase.Status == "error" {
		return c.JSON(http.StatusBadRequest, errBase)
	}

	return c.JSON(http.StatusOK, res)
}

func (api *BaseApi) DeleteDomain(c echo.Context) error {
	var baseError dto.BaseError

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		baseError.Status = "error"
		baseError.Message = "Invalid ID"
		return c.JSON(http.StatusBadRequest, baseError)
	}

	errBase := domainService.DeleteDomain(uint(id))
	if errBase != nil && errBase.Status == "error" {
		return c.JSON(http.StatusBadRequest, errBase)
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (api *BaseApi) GetListDomain(c echo.Context) error {
	res, errBase := domainService.GetListDomain()
	if errBase != nil && errBase.Status == "error" {
		return c.JSON(http.StatusBadRequest, errBase)
	}

	return c.JSON(http.StatusOK, res)
}

func (api *BaseApi) GetDomainByID(c echo.Context) error {
	var baseError dto.BaseError
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		baseError.Status = "error"
		baseError.Message = "Invalid ID"
		return c.JSON(http.StatusBadRequest, baseError)
	}

	res, errBase := domainService.GetDomainByID(uint(id))
	if errBase != nil && errBase.Status == "error" {
		return c.JSON(http.StatusBadRequest, errBase)
	}

	return c.JSON(http.StatusOK, res)
}
