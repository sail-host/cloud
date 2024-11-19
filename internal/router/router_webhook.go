package router

import (
	"github.com/labstack/echo/v4"
	apiV1 "github.com/sail-host/cloud/internal/app/api/v1"
)

type WebhookRouter struct {
}

func (r *WebhookRouter) InitRouter(Router *echo.Group) {
	webhookRouter := Router.Group("/webhook")

	webhookApi := apiV1.ApiGroupApp.BaseApi
	{
		webhookRouter.POST("/github/:project_id", webhookApi.WebhookGithub)
	}
}
