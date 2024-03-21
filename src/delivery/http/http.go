package http

import (
	"household-dashboard/src/controller"
	"household-dashboard/src/core/account"
	"household-dashboard/src/middleware"
	"household-dashboard/src/models"

	"github.com/gin-gonic/gin"
)

type Controllers struct {
	UserController    models.UserController
	AccountController account.AccountController
}

type HttpHandler interface {
	RegisterHttpHandler() *gin.Engine
}

func NewUserHttpHandler(controllers *controller.Controllers) HttpHandler {
	return &Controllers{
		UserController:    controllers.UserController,
		AccountController: controllers.AccountController,
	}
}

func (ctrl *Controllers) RegisterHttpHandler() *gin.Engine {
	r := gin.Default()

	r.GET(`/ping`, func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "PONG"})
	})

	r.POST(`/login`, ctrl.UserController.Login)
	r.POST(`/register`, ctrl.UserController.Register)

	acc := r.Group(`/account`)
	acc.Use(middleware.BearerTokenAuth())
	acc.GET("/id/:accountID", ctrl.AccountController.GetAccountByID)
	acc.GET("/me", ctrl.AccountController.GetAccountsByUserID)
	acc.POST("", ctrl.AccountController.CreateAccount)
	acc.GET("/types", ctrl.AccountController.GetAccountTypes)
	acc.POST("/types", ctrl.AccountController.CreateAccountType)
	acc.DELETE("/id/:accountID", ctrl.AccountController.DeleteAccount)

	return r
}
