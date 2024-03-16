package repository

import (
	"household-dashboard/src/models"
	accountrepo "household-dashboard/src/repository/account"
	userrepo "household-dashboard/src/repository/user"

	"github.com/jinzhu/gorm"
)

type Repositories struct {
	UserRepo    models.UserRepository
	AccountRepo models.AccountRepository
}

func InitRepositories(db *gorm.DB) *Repositories {
	userRepo := userrepo.NewUserRepository(db)
	accountRepo := accountrepo.NewAccountRepository(db)

	return &Repositories{
		UserRepo:    userRepo,
		AccountRepo: accountRepo,
	}
}
