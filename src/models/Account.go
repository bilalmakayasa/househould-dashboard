package models

import (
	"context"

	"github.com/gin-gonic/gin"
)

type AccountType struct {
	BaseModel
	Name string `json:"name"`
	Logo string `json:"logo"`
}

func (AccountType) TableName() string {
	return "account_type"
}

type AccountTypeInput struct {
	Name string `json:"name" binding:"required"`
	Logo string `json:"logo"`
}

type Account struct {
	BaseModel
	UserID        string  `json:"user_id"`
	AccountTypeID string  `json:"account_type_id"`
	Balance       float64 `json:"balance"`
	Currency      string  `json:"currency"`
	// create relation between UserID and User table
	User User `json:"user" gorm:"foreignkey:UserID"`
	// create relation between AccountTypeID and AccountType table
	AccountType AccountType `json:"account_type" gorm:"foreignkey:AccountTypeID"`
}

func (Account) TableName() string {
	return "account"
}

type AccountService interface {
	CreateAccountType(ctx context.Context, AccountTypeInput *AccountTypeInput) AccountType
	GetAccountTypes(ctx context.Context) []AccountType
}

type AccountRepository interface {
	CreateAccountType(AccountTypeInput *AccountTypeInput) AccountType
	// CreateAccount(account *Account) Account
	GetAccountTypes() []AccountType
	GetAccountByID(id string) Account
	GetAccountsByUserID(userID string) []Account
	// UpdateAccount(account *Account) Account
	// DeleteAccountByID(id string) bool
}

type AccountController interface {
	GetAccountTypes(ctx *gin.Context)
	CreateAccountType(ctx *gin.Context)
}
