package utils

import (
	"household-dashboard/src/config"
	"household-dashboard/src/models"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(tokenClaims *models.TokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "household-dashboard",
		},
		ID:   tokenClaims.ID,
		Name: tokenClaims.Name,
	})

	tokenString, err := token.SignedString([]byte(config.GetJwtKey()))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*models.TokenClaims, error) {
	claims := &models.TokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetJwtKey()), nil
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}

	return claims, nil
}
