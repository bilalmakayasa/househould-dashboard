package account

import (
	"household-dashboard/src/models"

	"github.com/gin-gonic/gin"
)

type AccountControllerHandler struct {
	service models.AccountService
}

func NewAccountControllerHandler(accountService models.AccountService) models.AccountController {
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
