package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sail-host/cloud/internal/app/dto"
)

func (b *BaseApi) CheckUpgrade(echo echo.Context) error {
	var baseError dto.BaseError

	res, err := upgradeService.CheckUpgrade(echo)
	if err != nil {
		baseError.Status = "error"
		baseError.Message = err.Error()
		return echo.JSON(http.StatusBadRequest, baseError)
	}

	return echo.JSON(http.StatusOK, res)
}
