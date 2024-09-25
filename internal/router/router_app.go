package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AppRouter struct {
}

func (r *AppRouter) InitRouter(Router *echo.Group) {
	// TODO: Remove this code after implementing the router
	Router.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
}
