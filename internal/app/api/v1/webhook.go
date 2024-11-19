package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sail-host/cloud/internal/app/dto"
	"github.com/sail-host/cloud/internal/global"
)

func (b *BaseApi) WebhookGithub(c echo.Context) error {
	var baseResponse dto.BaseResponse
	baseResponse.Status = "success"
	baseResponse.Message = "Webhook received"

	err := webhookGithubService.HandleWebhook(c)
	if err != nil {
		global.LOG.Error("Webhook error", err)
	}

	return c.JSON(http.StatusOK, baseResponse)
}
