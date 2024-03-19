package service

import (
	"household-dashboard/src/core/account"
	"household-dashboard/src/core/user"
	"household-dashboard/src/models"
	"household-dashboard/src/repository"
)

type Services struct {
	UserService    models.UserService
	AccountService models.AccountService
}

func InitServices(repo *repository.Repositories) *Services {
	userSerivce := user.InitUserService(repo.UserRepo)
	accountService := account.NewAccountService(repo.AccountRepo)

	return &Services{
		UserService:    userSerivce,
		AccountService: accountService,
	}
}
