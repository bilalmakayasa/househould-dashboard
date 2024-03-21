package models

type AccountType struct {
	BaseModel
	Name string `json:"name"`
	Logo string `json:"logo"`
}

func (AccountType) TableName() string {
	return "account_type"
}

type AccountTypeInput struct {
	Name string `json:"name" binding:"required"`
	Logo string `json:"logo"`
}

type Account struct {
	BaseModel
	UserID        string      `json:"user_id"`
	AccountTypeID string      `json:"account_type_id"`
	Balance       float64     `json:"balance"`
	Currency      string      `json:"currency"`
	User          User        `json:"user" gorm:"foreignkey:UserID"`
	AccountType   AccountType `json:"account_type" gorm:"foreignkey:AccountTypeID"`
}

func (Account) TableName() string {
	return "account"
}

type AccountInput struct {
	UserID        string  `json:"user_id"`
	AccountTypeID string  `json:"account_type_id" binding:"required"`
	Balance       float64 `json:"balance"`
	Currency      string  `json:"currency" binding:"required"`
}
