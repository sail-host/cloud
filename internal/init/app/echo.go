package app

import (
	"github.com/labstack/echo/v4"
	"github.com/sail-host/cloud/internal/global"
)

func Init() {
	global.ECHO = echo.New()
}
