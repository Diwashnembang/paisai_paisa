package main

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) home(c *gin.Context) {
	if value, exists := c.Keys["authorized"]; exists {
		if _, ok := value.(bool); !ok {
			c.Redirect(http.StatusSeeOther, "/login")
			slog.Info("no authorized : redirecting to login")
			return
		}
	} else {
		c.Redirect(http.StatusSeeOther, "/login")
		slog.Info("no authorized in the context")
		return

	}
	c.String(200, "hello world")
}

func (app *application) signUpPost(c *gin.Context) {
	//TODO : make third party sign in and up
	err := c.Request.ParseForm()
	if err != nil {
		c.String(http.StatusBadRequest, "Failed to parse form:%v ", err)
		return
	}
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")
	if email != "" && password != "" {

		_, err = app.DB.CreateUser(email, password)
		if err != nil {
			c.String(http.StatusInternalServerError, "error signing up")
			return
		}
	} else {
		c.String(http.StatusBadRequest, "email or passowrd cannot be empty")
		return
	}

	signature, err := app.generateJwtToken(email)
	if err != nil {
		slog.Error("error generating token")
		return

	}
	c.JSON(http.StatusAccepted, gin.H{
		"token": signature,
	})
}

func (app *application) loginPost(c *gin.Context) {

	err := c.Request.ParseForm()
	if err != nil {
		c.String(http.StatusBadRequest, "could not parse form:%s", err)
		return
	}
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")
	if email == "" && password == "" {
		c.String(http.StatusBadRequest, "email or passowrd cannot be empty")
		return
	}
	_, err = app.DB.VerifyUser(email, password)
	if err != nil {
		c.String(http.StatusInternalServerError, "error log in  ")
		slog.Info(err.Error())
		return
	}
	signature, err := app.generateJwtToken(email)
	if err != nil {
		slog.Error("error generating token")
	}
	c.JSON(http.StatusAccepted, gin.H{
		"token": signature,
	})

}
