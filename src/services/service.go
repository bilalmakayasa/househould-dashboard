package service

import userservice "household-dashboard/src/services/user"

type Services struct {
	UserService *userservice.UserService
}

type Repository struct {
	UserRepository *userservice.UserRepository
}

func InitServices(repo *Repository) *Services {
	userRepo := userservice.InitUserService(*repo.UserRepository)

	return &Services{
		UserService: userRepo,
	}
}
