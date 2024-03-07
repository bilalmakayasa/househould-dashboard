package user

import (
	"household-dashboard/src/models"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) CreateUser(user *models.RegisterPayload) models.User {
	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Username: user.Username,
		Phone:    user.Phone,
	}
	ur.db.Create(&newUser)
	return newUser
}

func (ur *UserRepository) GetUserByUsername(username string) models.User {
	var user models.User
	ur.db.Where("username = ?", username).First(&user)
	return user
}
