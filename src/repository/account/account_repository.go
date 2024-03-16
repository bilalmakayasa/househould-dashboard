package repository

import (
	"household-dashboard/src/models"
	"household-dashboard/src/utils"

	"github.com/jinzhu/gorm"
)

type AccountRepositoryHandler struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) models.AccountRepository {
	return &AccountRepositoryHandler{
		db: db,
	}
}

func (ar *AccountRepositoryHandler) GetAccountTypes() []models.AccountType {
	var accountTypes []models.AccountType
	ar.db.Find(&accountTypes)
	return accountTypes
}

func (ar *AccountRepositoryHandler) CreateAccountType(accountTypeInput *models.AccountTypeInput) models.AccountType {
	accountType := models.AccountType{
		BaseModel: models.BaseModel{
			ID: utils.GenerateUuid(),
		},
		Name: accountTypeInput.Name,
		Logo: accountTypeInput.Logo,
	}
	ar.db.Create(&accountType)
	return accountType
}

func (ar *AccountRepositoryHandler) GetAccountByID(id string) models.Account {
	var account models.Account
	ar.db.Where("id = ?", id).First(&account)
	return account
}

func (ar *AccountRepositoryHandler) GetAccountsByUserID(userID string) []models.Account {
	var accounts []models.Account
	ar.db.Where("user_id = ?", userID).Find(&accounts)
	return accounts
}
