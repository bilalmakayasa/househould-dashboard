package controller

import user "household-dashboard/src/controller/user"

type Controllers struct {
	UserController *user.UserController
}

type Services struct {
	UserService *user.UserService
}

func InitControllers(services *Services) *Controllers {
	userController := user.InitUserController(*services.UserService)

	return &Controllers{
		UserController: userController,
	}
}
