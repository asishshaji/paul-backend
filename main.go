package main

import (
	"backend/app"
	"backend/controller"
	"backend/repository"
	"backend/service"
	"backend/utils"
	"log"
)

func main() {
	db, err := utils.InitDB("mongodb://localhost:27017/myapp", "sada")
	if err != nil {
		log.Fatalln(err)
	}
	userRepo := repository.NewUserRepository(db, "users")
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	app := app.NewApp(":9091", userController)
	app.RunServer()
}
