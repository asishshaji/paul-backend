package app

import (
	"backend/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	e    *echo.Echo
	port string
}

func NewApp(port string, userController controller.IUserController) *App {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	uG := e.Group("/user")
	uG.POST("/", userController.CreateUser)

	return &App{
		e:    e,
		port: port,
	}
}

func (app *App) RunServer() {
	app.e.Logger.Fatal(app.e.Start(app.port))
}
