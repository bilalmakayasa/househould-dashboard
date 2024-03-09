package user

import "household-dashboard/src/models"

type UserServiceHandler struct {
	repo models.UserRepository
}

func InitUserService(repo models.UserRepository) models.UserService {
	return &UserServiceHandler{repo}
}

func (u *UserServiceHandler) Login(user models.LoginPayload) (models.User, error) {
	test := u.repo.GetUserByUsername(user.Username)

	return test, nil
}

func (u *UserServiceHandler) Register(user models.RegisterPayload) (models.User, error) {
	newUser := models.RegisterPayload{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Username: user.Username,
		Phone:    user.Phone,
	}
	createdUser := u.repo.CreateUser(&newUser)

	return createdUser, nil
}
