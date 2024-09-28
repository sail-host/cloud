package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sail-host/cloud/internal/app/dto"
)

func (api *BaseApi) Login(c echo.Context) error {
	request := dto.LoginRequest{}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	res, err := authService.Login(c, request)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	return c.JSON(http.StatusOK, res)
}

func (api *BaseApi) CheckUserFirstTime(c echo.Context) error {
	exist, err := authService.CheckUserFirstTime(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	var response dto.BaseResponse

	if exist {
		response.Status = "error"
		response.Message = "User already exists"
	} else {
		response.Status = "success"
		response.Message = "User first time"
	}

	return c.JSON(http.StatusOK, response)
}

func (api *BaseApi) Register(c echo.Context) error {
	request := dto.RegisterRequest{}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	res, err := authService.Register(c, request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, res)
}
