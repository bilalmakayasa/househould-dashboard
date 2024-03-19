package repository

import (
	account "household-dashboard/src/core/account"
	"household-dashboard/src/core/user"
	"household-dashboard/src/models"

	"github.com/jinzhu/gorm"
)

type Repositories struct {
	UserRepo    models.UserRepository
	AccountRepo models.AccountRepository
}

func InitRepositories(db *gorm.DB) *Repositories {
	userRepo := user.NewUserRepository(db)
	accountRepo := account.NewAccountRepository(db)

	return &Repositories{
		UserRepo:    userRepo,
		AccountRepo: accountRepo,
	}
}
