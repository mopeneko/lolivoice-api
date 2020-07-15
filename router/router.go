package router

import (
	"github.com/labstack/echo/v4"
	"github.com/mopeneko/lolivoice-api/controller"
)

var Router *echo.Echo

func init() {
	Router = echo.New()

	Router.GET("/voice", controller.GetVoice)
}

func Run() {
	Router.Logger.Fatal(Router.Start(":1323"))
}

