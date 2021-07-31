package controller

import (
	"backend/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type QuizController struct {
	quizService service.IQuizService
}

func NewQuizController(quizService service.IQuizService) IQuizController {
	return &QuizController{
		quizService: quizService,
	}
}

func (controller QuizController) StartQuiz(c echo.Context) error {

	genre := "comedy"

	quizArray := controller.quizService.GetQuizBasedOnGenre(c.Request().Context(), genre)

	// Return 5 questions based on this genre

	return c.JSON(http.StatusOK, quizArray)
}
