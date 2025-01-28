package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (app *application) generateJwtToken(email string) (string, error) {

	key := os.Getenv("JWT_TOKEN")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": email,
		"exp":  time.Now().Add(24 * time.Hour),
	})
	signature, err := token.SignedString([]byte(key))
	if err != nil {

		slog.Error(fmt.Sprintf("error signing the key: %s", err))
		return "", errors.New("erroe signing the key")
	}
	return signature, nil

}
