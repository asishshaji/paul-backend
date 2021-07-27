package controller

import (
	"backend/dto"
	"backend/service"
	"log"
	"net/http"

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
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, "Created user successfully")
}

func (uC UserController) LoginUser(ctx echo.Context) error {
	return nil
}
