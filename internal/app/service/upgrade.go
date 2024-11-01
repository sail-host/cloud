package service

import (
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sail-host/cloud/internal/app/dto"
	"github.com/sail-host/cloud/internal/constants"
	"github.com/sail-host/cloud/internal/global"
	"github.com/sail-host/cloud/internal/utils/sailhost"
)

type UpgradeService struct{}

type IUpgradeService interface {
	CheckUpgrade(c echo.Context) (*dto.BaseResponse, error)
	Update(c echo.Context) (*dto.BaseResponse, error)
}

func NewIUpgradeService() IUpgradeService {
	return &UpgradeService{}
}

func (u *UpgradeService) CheckUpgrade(c echo.Context) (*dto.BaseResponse, error) {
	lastVersion, err := sailhost.LastVersion()
	if err != nil {
		return nil, err
	}

	var baseResponse dto.BaseResponse
	baseResponse.Data = map[string]string{
		"last_version":    lastVersion,
		"current_version": constants.Version,
	}
	baseResponse.Message = "Check upgrade success"
	baseResponse.Status = "success"

	return &baseResponse, nil
}

func (u *UpgradeService) Update(c echo.Context) (*dto.BaseResponse, error) {
	var baseResponse dto.BaseResponse

	err := sailhost.DownloadLastVersion(global.CONF.System.BaseDir)
	if err != nil {
		return nil, err
	}

	go restart()

	baseResponse.Message = "Update success"
	baseResponse.Status = "success"

	return &baseResponse, nil
}

func restart() {
	time.Sleep(1 * time.Second)
	os.Exit(0)
}
