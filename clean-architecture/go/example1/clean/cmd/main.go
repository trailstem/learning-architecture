package main

import (
	"fmt"
	"myapp/controllers"
	"myapp/repositories"
)

func main() {
	userRepo := repositories.NewUserRepository()
	userController := controllers.NewUserController(userRepo)

	user, err := userController.GetUser(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user)
	}
}
