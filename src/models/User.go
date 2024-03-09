package models

import "github.com/gin-gonic/gin"

type User struct {
	BaseModel
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
}

func (User) TableName() string {
	return "user"
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
}

type UserRepository interface {
	CreateUser(user *RegisterPayload) User
	GetUserByUsername(username string) User
}

type UserService interface {
	Login(user LoginPayload) (User, error)
	Register(user RegisterPayload) (User, error)
}

type UserController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}
