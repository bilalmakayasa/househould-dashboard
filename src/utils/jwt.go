package utils

import (
	"household-dashboard/src/config"
	"household-dashboard/src/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(tokenClaims *models.TokenClaims) (string, error) {
	claims := &models.TokenClaims{
		ID:   tokenClaims.ID,
		Name: tokenClaims.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.GetJwtKey()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(token string) (*models.TokenClaims, error) {
	claims := &models.TokenClaims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetJwtKey()), nil
	})
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, err
	}

	return claims, nil
}
