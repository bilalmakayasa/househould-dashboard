package controller

import (
	"household-dashboard/src/models"

	"github.com/gin-gonic/gin"
)

type UserControllerHandler struct {
	service models.UserService
}

func InitUserController(userService models.UserService) models.UserController {
	return &UserControllerHandler{userService}
}

func (u *UserControllerHandler) Login(c *gin.Context) {
	var user models.LoginPayload
	c.BindJSON(&user)

	if user.Username == "" || user.Password == "" {
		c.JSON(400, gin.H{"error": "Username and password are required"})
		return
	}

	userService, err := u.service.Login(user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		panic(err)
	}
	c.JSON(200, userService)
}

func (u *UserControllerHandler) Register(c *gin.Context) {
	var user models.RegisterPayload
	c.BindJSON(&user)

	if user.Name == "" || user.Email == "" || user.Password == "" || user.Username == "" || user.Phone == "" {
		c.JSON(400, gin.H{"error": "All fields are required"})
		return
	}

	newUser, err := u.service.Register(user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, newUser)
}
