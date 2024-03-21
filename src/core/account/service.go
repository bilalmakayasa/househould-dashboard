package account

import (
	"context"
	"errors"
	"household-dashboard/src/models"
	"household-dashboard/src/utils"
)

type AccountService interface {
	CreateAccountType(ctx context.Context, AccountTypeInput *models.AccountTypeInput) models.AccountType
	GetAccountTypes(ctx context.Context) []models.AccountType
	CreateAccount(ctx context.Context, accountInput *models.AccountInput) (models.Account, error)
	GetAccountByID(ctx context.Context, accountID string) models.Account
	GetAccountsByUserID(ctx context.Context) []models.Account
	DeleteAccount(ctx context.Context, accountID string) bool
}

type AccountServiceHandler struct {
	repo AccountRepository
}

func NewAccountService(repo AccountRepository) AccountService {
	return &AccountServiceHandler{repo}
}

func (a *AccountServiceHandler) GetAccountTypes(ctx context.Context) []models.AccountType {
	return a.repo.GetAccountTypes()
}

func (a *AccountServiceHandler) CreateAccountType(ctx context.Context, accountTypeInput *models.AccountTypeInput) models.AccountType {
	return a.repo.CreateAccountType(accountTypeInput)
}

func (a *AccountServiceHandler) CreateAccount(ctx context.Context, accountInput *models.AccountInput) (models.Account, error) {
	context := utils.GetContextValue(ctx)

	if context.User.ID == "" {
		return models.Account{}, errors.New("user not found")
	}

	accountType := a.repo.GetAccountTypeByID(accountInput.AccountTypeID)

	if accountType.ID == "" {
		return models.Account{}, errors.New("account type not found")
	}

	validateAccount := a.repo.GetAccountByIDandAccountType(context.User.ID, accountInput.AccountTypeID)

	if validateAccount.BaseModel.ID != "" {
		return models.Account{}, errors.New("account already exists")
	}

	accountPayload := models.AccountInput{
		UserID:        context.User.ID,
		AccountTypeID: accountInput.AccountTypeID,
		Balance:       accountInput.Balance,
		Currency:      accountInput.Currency,
	}

	test := a.repo.CreateAccount(&accountPayload)

	return test, nil
}

func (a *AccountServiceHandler) GetAccountByID(ctx context.Context, accountID string) models.Account {
	return a.repo.GetAccountByID(accountID)
}

func (a *AccountServiceHandler) GetAccountsByUserID(ctx context.Context) []models.Account {
	context := utils.GetContextValue(ctx)

	return a.repo.GetAccountsByUserID(context.User.ID)
}

func (a *AccountServiceHandler) DeleteAccount(ctx context.Context, accountID string) bool {
	return a.repo.DeleteAccount(accountID)
}
