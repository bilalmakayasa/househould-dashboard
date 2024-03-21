package account

import (
	"fmt"
	"household-dashboard/src/models"
	"household-dashboard/src/utils"

	"github.com/jinzhu/gorm"
)

type AccountRepository interface {
	CreateAccountType(AccountTypeInput *models.AccountTypeInput) models.AccountType
	CreateAccount(account *models.AccountInput) models.Account
	GetAccountTypes() []models.AccountType
	GetAccountTypeByID(id string) models.AccountType
	GetAccountByID(id string) models.Account
	GetAccountsByUserID(userID string) []models.Account
	UpdateAccount(account *models.Account) models.Account
	DeleteAccountByID(id string) bool
	GetAccountByIDandAccountType(id string, accountTypeId string) models.Account
	DeleteAccount(id string) bool
}

type AccountRepositoryHandler struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &AccountRepositoryHandler{
		db: db,
	}
}

func (ar *AccountRepositoryHandler) GetAccountTypes() []models.AccountType {
	var accountTypes []models.AccountType
	ar.db.Find(&accountTypes)
	return accountTypes
}

func (ar *AccountRepositoryHandler) GetAccountTypeByID(id string) models.AccountType {
	var accountType models.AccountType
	ar.db.Where("id = ?", id).First(&accountType)
	return accountType
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
	err := ar.db.Where("id = ?", id).First(&account).Error
	if err != nil {
		fmt.Println(err)
		return models.Account{}
	}
	return account
}

func (ar *AccountRepositoryHandler) GetAccountsByUserID(userID string) []models.Account {
	var accounts []models.Account

	ar.db.Where("user_id = ?", userID).Find(&accounts)
	return accounts
}

func (ar *AccountRepositoryHandler) CreateAccount(accountInput *models.AccountInput) models.Account {
	account := models.Account{
		BaseModel: models.BaseModel{
			ID: utils.GenerateUuid(),
		},
		UserID:        accountInput.UserID,
		AccountTypeID: accountInput.AccountTypeID,
		Balance:       accountInput.Balance,
		Currency:      accountInput.Currency,
	}

	ar.db.Create(&account)
	return account
}

func (ar *AccountRepositoryHandler) UpdateAccount(account *models.Account) models.Account {
	ar.db.Save(&account)
	return *account
}

func (ar *AccountRepositoryHandler) DeleteAccountByID(id string) bool {
	ar.db.Where("id = ?", id).Delete(&models.Account{})
	return true
}

func (ar *AccountRepositoryHandler) GetAccountByIDandAccountType(userId string, accountTypeId string) models.Account {
	var account models.Account

	ar.db.Where("user_id = ? AND account_type_id = ?", userId, accountTypeId).First(&account)

	return account
}

func (ar *AccountRepositoryHandler) DeleteAccount(id string) bool {
	var account models.Account

	ar.db.Where("id = ?", id).Delete(&account)

	return true
}
