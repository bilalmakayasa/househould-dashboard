package user

import "household-dashboard/src/models"

type UserRepository interface {
	CreateUser(user *models.User) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
}

type UserService struct {
	repo UserRepository
}

func InitUserService(repo UserRepository) *UserService {
	return &UserService{repo}
}
