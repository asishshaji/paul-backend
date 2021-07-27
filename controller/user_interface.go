package controller

import "github.com/labstack/echo/v4"

type IUserController interface {
	CreateUser(c echo.Context) error
	LoginUser(c echo.Context) error
	// UpdateUser(c echo.Context) error
}
