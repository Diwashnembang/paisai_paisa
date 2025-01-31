package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type jsonResponse struct {
	Success bool           `json:"success"`
	Payload map[string]any `json:"payload"`
}

func (app *application) generateJwtToken(email string) (string, error) {

	key := os.Getenv("JWT_TOKEN")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": email,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	})
	signature, err := token.SignedString([]byte(key))
	if err != nil {

		slog.Error(fmt.Sprintf("error signing the key: %s", err))
		return "", errors.New("erroe signing the key")
	}
	return signature, nil

}

// templete to send error response in json
func (app *application) errorJson(errMsg string) *jsonResponse {
	return &jsonResponse{
		Success: false,
		Payload: gin.H{
			"error": errMsg,
		},
	}
}

// templete to send sucess response in json
func (app *application) successJson(token string) *jsonResponse {
	return &jsonResponse{
		Success: true,
		Payload: gin.H{
			"token": token,
		},
	}
}
