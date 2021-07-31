package controller

import "github.com/labstack/echo/v4"

type IQuizController interface {
	StartQuiz(c echo.Context) error
}
