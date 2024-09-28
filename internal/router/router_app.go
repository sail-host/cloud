package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sail-host/cloud/internal/app/dto"
)

type AppRouter struct {
}

func (r *AppRouter) InitRouter(Router *echo.Group) {
	// TODO: Remove this code after implementing the router
	Router.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, dto.BaseResponse{Status: "success", Message: "pong"})
	})
}
