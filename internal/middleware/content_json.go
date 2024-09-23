package middleware

import (
	"github.com/labstack/echo/v4"
)

// ContentJSON is a middleware that sets the content type to json
func ContentJSON(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "application/json")
		return next(c)
	}
}
