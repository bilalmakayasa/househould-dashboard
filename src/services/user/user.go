package services

import (
	"errors"
	"household-dashboard/src/models"
	"household-dashboard/src/utils"
)

type UserServiceHandler struct {
	repo models.UserRepository
}

func InitUserService(repo models.UserRepository) models.UserService {
	return &UserServiceHandler{repo}
}

func (u *UserServiceHandler) Login(user models.LoginPayload) (models.LoginResponse, error) {
	test := u.repo.GetUserByUsername(user.Username)

	if test.ID == "" {
		return models.LoginResponse{}, errors.New("user not found")
	}

	err := utils.VerifyPassword(test.Password, user.Password)
	if err != nil {
		return models.LoginResponse{}, errors.New(err.Error())
	}

	tokenClaims, err := utils.GenerateToken(&models.TokenClaims{
		ID:   test.ID,
		Name: test.Name,
	})
	if err != nil {
		return models.LoginResponse{}, err
	}

	result := &models.LoginResponse{
		User: models.User{
			BaseModel: models.BaseModel{
				ID: test.ID,
			},
			Name:  test.Name,
			Email: test.Email,
			Phone: test.Phone,
		},
		Token: tokenClaims,
	}

	return *result, nil
}

func (u *UserServiceHandler) Register(user models.RegisterPayload) (models.User, error) {

	hashedPassword, err := utils.CreatePassword(user.Password)

	if err != nil {
		return models.User{}, err
	}

	newUser := models.RegisterPayload{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
		Username: user.Username,
		Phone:    user.Phone,
	}
	createdUser := u.repo.CreateUser(&newUser)

	return createdUser, nil
}
