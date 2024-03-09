package repository

import (
	"household-dashboard/src/models"
	userrepo "household-dashboard/src/repository/user"

	"github.com/jinzhu/gorm"
)

type Repositories struct {
	UserRepo models.UserRepository
}

func InitRepositories(db *gorm.DB) *Repositories {
	userRepo := userrepo.NewUserRepository(db)

	return &Repositories{
		UserRepo: userRepo,
	}
}
