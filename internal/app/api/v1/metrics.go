package v1

import (
	"github.com/labstack/echo/v4"
)

func (s *BaseApi) GetSystemMetrics(c echo.Context) error {
	metrics, err := metricsService.GetSystemMetrics()
	if err != nil {
		return c.JSON(500, echo.Map{"error": err.Error()})
	}
	return c.JSON(200, metrics)
}
