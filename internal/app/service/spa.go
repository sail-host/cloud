package service

import (
	"io/fs"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type SPA struct {
	assets fs.FS
}

func NewSPA(assets fs.FS) *SPA {
	return &SPA{assets: assets}
}

func (s *SPA) ServeHTTP() echo.HandlerFunc {
	return func(c echo.Context) error {
		url := c.Request().URL.Path
		if !strings.Contains(url, ".") {
			c.Request().URL.Path = "/"
		}
		http.FileServer(http.FS(s.assets)).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}
}
