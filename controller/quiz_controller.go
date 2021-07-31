package controller

import (
	"backend/service"
	"log"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
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
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims) // claims
	log.Println(claims)

	quizArray := controller.quizService.GetQuizBasedOnGenre(c.Request().Context(), genre)

	return c.JSON(http.StatusOK, quizArray)

}

func (controller QuizController) SaveScore(c echo.Context) error {
	score := c.FormValue("score")
	genre := c.FormValue("genre")

	scoreInt, err := strconv.Atoi(score)
	if err != nil {
		return echo.ErrInternalServerError
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims) // claims
	log.Println(claims)

	username := claims["username"].(string)

	err = controller.quizService.SaveScore(c.Request().Context(), scoreInt, username, genre)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{username: score})
}
