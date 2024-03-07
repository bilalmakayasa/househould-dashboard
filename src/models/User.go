package models

type User struct {
	BaseModel
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
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
