package controller

import (
	user "household-dashboard/src/controller/user"
	"household-dashboard/src/models"
	service "household-dashboard/src/services"
)

type Controllers struct {
	UserController models.UserController
}

func InitControllers(services *service.Services) *Controllers {
	userController := user.InitUserController(services.UserService)

	return &Controllers{
		UserController: userController,
	}
}
