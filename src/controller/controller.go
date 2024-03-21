package controller

import (
	"household-dashboard/src/core/account"
	"household-dashboard/src/core/user"
	"household-dashboard/src/models"
	service "household-dashboard/src/services"
)

type Controllers struct {
	UserController    models.UserController
	AccountController account.AccountController
}

func InitControllers(services *service.Services) *Controllers {
	userController := user.InitUserController(services.UserService)
	accountController := account.NewAccountControllerHandler(services.AccountService)

	return &Controllers{
		UserController:    userController,
		AccountController: accountController,
	}
}
