package models

import (
	"time"

	"github.com/google/uuid"
)

type UUID uuid.UUID

type BaseModel struct {
	ID        string    `primary_key;" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Tabler interface {
	TableName() string
}
