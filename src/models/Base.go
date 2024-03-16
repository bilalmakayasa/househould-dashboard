package models

import (
	"time"

	"github.com/google/uuid"
)

type UUID uuid.UUID

type BaseModel struct {
	ID        string    `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ContextType struct {
	User TokenClaims `json:"user"`
}
