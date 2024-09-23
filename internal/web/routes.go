package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Route struct {
	echo *echo.Echo
}

func NewRoute(e *echo.Echo) *Route {
	return &Route{echo: e}
}

func (r *Route) Run() {
	api := r.echo.Group("/api")
	// api.Use(middleware.ContentJSON)

	// Ping
	api.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "{\"message\":\"pong\"}")
	})
}
