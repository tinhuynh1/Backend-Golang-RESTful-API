package router

import (
	"DACN-GithubTrending/handler"
	myMiddleware "DACN-GithubTrending/middleware"
	"github.com/labstack/echo"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter() {
	api.Echo.POST("/user/sign-in", api.UserHandler.HandleSignIn, myMiddleware.ISAdmin())
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)
}
