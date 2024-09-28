package router

import "github.com/labstack/echo/v4"

type ICommonRouter interface {
	InitRouter(Router *echo.Group)
}

func commonGroups() []ICommonRouter {
	return []ICommonRouter{
		&BaseRouter{},
		&AppRouter{},
		&UserRouter{},
	}
}
