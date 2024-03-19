package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func CreatePassword(str string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func VerifyPassword(hashedPassword, password string) error {
	fmt.Println(hashedPassword, password)
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
