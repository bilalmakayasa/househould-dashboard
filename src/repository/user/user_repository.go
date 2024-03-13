package repository

import (
	"household-dashboard/src/models"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type UserRepositoryHandler struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) models.UserRepository {
	return &UserRepositoryHandler{
		db: db,
	}
}

func (ur *UserRepositoryHandler) CreateUser(user *models.RegisterPayload) models.User {
	newUser := models.User{
		BaseModel: models.BaseModel{
			ID: uuid.NewString(),
		},
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Username: user.Username,
		Phone:    user.Phone,
	}
	ur.db.Create(&newUser)
	return newUser
}

func (ur *UserRepositoryHandler) GetUserByUsername(username string) models.User {
	var user models.User
	ur.db.Where("username = ?", username).First(&user)
	return user
}
