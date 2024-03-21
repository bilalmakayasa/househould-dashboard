package account

import (
	"household-dashboard/src/models"

	"github.com/gin-gonic/gin"
)

type AccountController interface {
	GetAccountTypes(ctx *gin.Context)
	CreateAccountType(ctx *gin.Context)
	CreateAccount(ctx *gin.Context)
	GetAccountByID(ctx *gin.Context)
	GetAccountsByUserID(ctx *gin.Context)
	DeleteAccount(ctx *gin.Context)
}

type AccountControllerHandler struct {
	service AccountService
}

func NewAccountControllerHandler(accountService AccountService) AccountController {
	return &AccountControllerHandler{accountService}
}

func (ac *AccountControllerHandler) GetAccountTypes(c *gin.Context) {
	accountTypes := ac.service.GetAccountTypes(c)

	c.JSON(200, gin.H{"data": accountTypes})
}

func (ac *AccountControllerHandler) CreateAccountType(c *gin.Context) {
	var accountTypeInput models.AccountTypeInput
	if err := c.ShouldBindJSON(&accountTypeInput); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	accountType := ac.service.CreateAccountType(c, &accountTypeInput)

	c.JSON(200, gin.H{"data": accountType})
}

func (ac *AccountControllerHandler) CreateAccount(c *gin.Context) {
	var accountInput models.AccountInput
	if err := c.ShouldBindJSON(&accountInput); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	account, err := ac.service.CreateAccount(c, &accountInput)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": account})
}

func (ac *AccountControllerHandler) GetAccountByID(c *gin.Context) {
	accountID := c.Param("accountID")

	account := ac.service.GetAccountByID(c, accountID)

	c.JSON(200, gin.H{"data": account})
}

func (ac *AccountControllerHandler) GetAccountsByUserID(c *gin.Context) {

	accounts := ac.service.GetAccountsByUserID(c)

	c.JSON(200, gin.H{"data": accounts})
}

func (ac *AccountControllerHandler) DeleteAccount(c *gin.Context) {
	accountID := c.Param("accountID")

	account := ac.service.DeleteAccount(c, accountID)
	c.JSON(200, gin.H{"data": account})
}
