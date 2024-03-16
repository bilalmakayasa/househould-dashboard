package service

import (
	"household-dashboard/src/models"
	"household-dashboard/src/repository"
	accountService "household-dashboard/src/services/account"
	userservice "household-dashboard/src/services/user"
)

type Services struct {
	UserService    models.UserService
	AccountService models.AccountService
}

func InitServices(repo *repository.Repositories) *Services {
	userSerivce := userservice.InitUserService(repo.UserRepo)
	accountService := accountService.NewAccountService(repo.AccountRepo)

	return &Services{
		UserService:    userSerivce,
		AccountService: accountService,
	}
}
