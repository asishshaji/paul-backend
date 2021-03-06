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

func NewApp(port string, userController controller.IUserController, quizController controller.IQuizController) *App {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	uG := e.Group("/user")
	uG.POST("/", userController.CreateUser)
	uG.POST("/authenticate", userController.LoginUser)
	uG.Use(middleware.JWT([]byte("secret")))
	uG.POST("/genre", userController.AddGenre)

	qG := e.Group("/quiz")
	qG.Use(middleware.JWT([]byte("secret")))
	qG.GET("/start", quizController.StartQuiz)
	qG.POST("/", quizController.SaveScore)

	return &App{
		e:    e,
		port: port,
	}
}

func (app *App) RunServer() {
	app.e.Logger.Fatal(app.e.Start(app.port))
}
