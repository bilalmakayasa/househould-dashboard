package user

import (
	"household-dashboard/src/models"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Login(user models.LoginPayload) (models.User, error)
	Register(user models.RegisterPayload) (models.User, error)
}

type UserController struct {
	service UserService
}

func InitUserController(userService UserService) *UserController {
	return &UserController{userService}
}

func (u *UserController) Login(c *gin.Context) error {
	var user models.LoginPayload
	c.BindJSON(&user)

	if user.Username == "" || user.Password == "" {
		c.JSON(400, gin.H{"error": "Username and password are required"})
		return nil
	}

	test, err := u.service.Login(user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid username or password"})
		return nil
	}
	c.JSON(200, test)
	return nil
}
