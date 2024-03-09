package service

import (
	"household-dashboard/src/models"
	"household-dashboard/src/repository"
	userservice "household-dashboard/src/services/user"
)

type Services struct {
	UserService models.UserService
}

func InitServices(repo *repository.Repositories) *Services {
	userSerivce := userservice.InitUserService(repo.UserRepo)

	return &Services{
		UserService: userSerivce,
	}
}
