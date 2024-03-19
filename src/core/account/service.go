package account

import (
	"context"
	"household-dashboard/src/models"
)

type AccountServiceHandler struct {
	repo models.AccountRepository
}

func NewAccountService(repo models.AccountRepository) models.AccountService {
	return &AccountServiceHandler{repo}
}

func (a *AccountServiceHandler) GetAccountTypes(ctx context.Context) []models.AccountType {
	return a.repo.GetAccountTypes()
}

func (a *AccountServiceHandler) CreateAccountType(ctx context.Context, accountTypeInput *models.AccountTypeInput) models.AccountType {
	return a.repo.CreateAccountType(accountTypeInput)
}
