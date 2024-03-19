package http

import (
	"household-dashboard/src/controller"
	"household-dashboard/src/middleware"
	"household-dashboard/src/models"

	"github.com/gin-gonic/gin"
)

type Controllers struct {
	UserController    models.UserController
	AccountController models.AccountController
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
	acc.GET("/types", ctrl.AccountController.GetAccountTypes)
	acc.POST("/types", ctrl.AccountController.CreateAccountType)

	return r
}
