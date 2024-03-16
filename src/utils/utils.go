package utils

import (
	"context"
	"household-dashboard/src/models"

	"github.com/google/uuid"
)

func GenerateUuid() string {
	return uuid.New().String()
}

func GetContextValue(ctx context.Context) models.ContextType {
	user := ctx.Value("user").(*models.TokenClaims)
	result := &models.ContextType{
		User: *user,
	}
	return *result
}
