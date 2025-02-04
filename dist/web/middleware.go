package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
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
	authHear := c.GetHeader("Authorization")

	if !strings.HasPrefix(authHear, "Bearer") {
		c.String(http.StatusBadRequest, "Not Authorized no bearer")
		c.Abort()
		return
	}
	rawToken := strings.TrimSpace(strings.TrimPrefix(authHear, "Bearer"))
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
	var u string
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		u, err = claims.GetSubject()
		if err != nil {
			slog.Error("couldnot get sub", err.Error())
			return
		}

	} else {
		slog.Error("unknown claims type, cannot proceed")
		return
	}
	userID, err := strconv.ParseUint(u, 10, 32)
	if err != nil {
		slog.Error("error converiting userid into int", err.Error())
		c.JSON(http.StatusInternalServerError, app.errorJson("error verifying user"))
	}
	c.Set("userID", uint(userID))
	c.Next()

}
