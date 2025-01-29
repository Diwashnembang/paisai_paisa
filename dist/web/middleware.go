package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (app *application) allowedCors() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	return cors.New(corsConfig)
}

func (app *application) isAuthorized(c *gin.Context) {
	//TODO : get claims form the token
	authHear := c.GetHeader("Authorization")

	if !strings.HasPrefix(authHear, "Bearer") {
		c.String(http.StatusBadRequest, "Not Authorized no bearer")
		c.Abort()
		return
	}
	rawToken := strings.TrimSpace(strings.TrimPrefix(authHear, "Bearer"))
	slog.Info(rawToken)
	token, err := jwt.Parse(rawToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected sigining method; %v", t.Header["alg"])
		}
		return []byte(os.Getenv("JWT_TOKEN")), nil
	})
	if err != nil {
		c.String(http.StatusBadRequest, "Not Authorized")
		c.Abort()
		slog.Info("unauthorized access attemp", err)
		return
	}

	token.Valid = true
	c.Next()

}
