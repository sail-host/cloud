package service

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sail-host/cloud/internal/app/dto"
	"github.com/sail-host/cloud/internal/config"
)

const (
	UpgradeUrl = "https://api.github.com/repos/sail-host/cloud/tags"
)

type UpgradeService struct{}

type IUpgradeService interface {
	CheckUpgrade(c echo.Context) (*dto.BaseResponse, error)
}

func NewIUpgradeService() IUpgradeService {
	return &UpgradeService{}
}

func (u *UpgradeService) CheckUpgrade(c echo.Context) (*dto.BaseResponse, error) {
	type response struct {
		Name       string `json:"name"`
		TarballUrl string `json:"tarball_url"`
		Commit     struct {
			Sha string `json:"sha"`
			Url string `json:"url"`
		} `json:"commit"`
		NodeId string `json:"node_id"`
	}

	var resp []response

	res, err := http.Get(UpgradeUrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	var baseResponse dto.BaseResponse
	baseResponse.Data = map[string]string{
		"last_version":    resp[0].Name,
		"current_version": config.Version,
	}
	baseResponse.Message = "Check upgrade success"
	baseResponse.Status = "success"

	return &baseResponse, nil
}
