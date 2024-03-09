package http

import (
	"household-dashboard/src/controller"
	"household-dashboard/src/models"

	"github.com/gin-gonic/gin"
)

type Controllers struct {
	UserController models.UserController
}

type HttpHandler interface {
	RegisterHttpHandler() *gin.Engine
}

func NewUserHttpHandler(controllers *controller.Controllers) HttpHandler {
	return &Controllers{
		UserController: controllers.UserController,
	}
}

func (ctrl *Controllers) RegisterHttpHandler() *gin.Engine {
	r := gin.Default()

	r.GET(`/ping`, func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "PONG"})
	})

	r.POST(`/login`, ctrl.UserController.Login)
	r.POST(`/register`, ctrl.UserController.Register)

	return r
}
