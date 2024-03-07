package user

import "household-dashboard/src/models"

func (u *UserService) Login(user models.LoginPayload) (models.User, error) {
	test, err := u.repo.GetUserByUsername(user.Username)
	if err != nil {
		return models.User{}, err
	}

	return test, nil
}

func (u *UserService) Register(user models.RegisterPayload) (models.User, error) {
	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Username: user.Username,
		Phone:    user.Phone,
	}
	createdUser, err := u.repo.CreateUser(&newUser)
	if err != nil {
		return models.User{}, err
	}

	return createdUser, nil
}
