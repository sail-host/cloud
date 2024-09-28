package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sail-host/cloud/internal/app/dto"
	"github.com/sail-host/cloud/internal/app/model"
)

func (api *BaseApi) GetUser(c echo.Context) error {
	var response dto.BaseResponse
	user := c.Get("user").(*model.User)
	user.Password = ""

	response.Status = "success"
	response.Message = "User fetched successfully"
	response.Data = user

	return c.JSON(http.StatusOK, response)
}
