package main

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) home(c *gin.Context) {
	c.String(200, "hello world")
}

func (app *application) signUpPost(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.String(http.StatusBadRequest, "Failed to parse form:%v ", err)
		return
	}
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")
	_, err = app.DB.CreateUser(email, password)
	if err != nil {
		c.String(http.StatusInternalServerError, "error signing up")
		return
	}
	// c.Redirect(http.StatusSeeOther, "/")
	signature, err := app.generateJwtToken(email)
	if err != nil {
		slog.Error("error generating token")
		return

	}
	c.String(http.StatusAccepted, signature)
	slog.Info("still running")

}
