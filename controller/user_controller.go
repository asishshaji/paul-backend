package controller

import (
	"backend/dto"
	"backend/service"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(uS service.IUserService) IUserController {
	return UserController{
		userService: uS,
	}
}

func (uC UserController) CreateUser(ctx echo.Context) error {

	userReq := new(dto.UserRegDto)
	if err := ctx.Bind(userReq); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	user := dto.UserRegDto{
		Username: userReq.Username,
		Password: userReq.Password,
	}

	err := uC.userService.CreateUser(ctx.Request().Context(), user)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, "User already exists")
	}

	return ctx.JSON(http.StatusOK, "Created user successfully")
}

func (uC UserController) LoginUser(ctx echo.Context) error {
	userReq := new(dto.UserRegDto)
	if err := ctx.Bind(userReq); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	user := dto.UserRegDto{
		Username: userReq.Username,
		Password: userReq.Password,
	}

	err := uC.userService.LoginUser(ctx.Request().Context(), user)

	if err != nil {
		log.Println(err)
		return echo.ErrUnauthorized
	}
	if err != nil {
		log.Println(err)
		return echo.ErrInternalServerError
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Println(err)
		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"token": t,
	})

}
