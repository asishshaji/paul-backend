package controller

import "github.com/labstack/echo/v4"

type IUserController interface {
	CreateUser(c echo.Context) error
	LoginUser(c echo.Context) error
	AddGenre(c echo.Context) error
}
