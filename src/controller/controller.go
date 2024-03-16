package controller

import (
	account "household-dashboard/src/controller/account"
	user "household-dashboard/src/controller/user"
	"household-dashboard/src/models"
	service "household-dashboard/src/services"
)

type Controllers struct {
	UserController    models.UserController
	AccountController models.AccountController
}

func InitControllers(services *service.Services) *Controllers {
	userController := user.InitUserController(services.UserService)
	accountController := account.NewAccountControllerHandler(services.AccountService)

	return &Controllers{
		UserController:    userController,
		AccountController: accountController,
	}
}
