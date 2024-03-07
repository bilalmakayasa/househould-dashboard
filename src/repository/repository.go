package repository

import (
	userrepo "household-dashboard/src/repository/user"

	"github.com/jinzhu/gorm"
)

type Repositories struct {
	UserRepo *userrepo.UserRepository
}

func InitRepositories(db *gorm.DB) *Repositories {
	userRepo := userrepo.NewUserRepository(db)
	return &Repositories{
		UserRepo: userRepo,
	}
}
