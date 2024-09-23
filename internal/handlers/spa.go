package handlers

import (
	"io/fs"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func SPA(assets fs.FS) echo.HandlerFunc {
	return func(c echo.Context) error {
		url := c.Request().URL.Path
		if !strings.Contains(url, ".") {
			c.Request().URL.Path = "/"
		}
		http.FileServer(http.FS(assets)).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}
}
