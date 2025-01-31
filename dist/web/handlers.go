package main

import (
	"fmt"
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
		c.JSON(http.StatusBadRequest, app.errorJson(fmt.Sprintf("Failed to parse form : %s", err)))
		return
	}
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")
	if email != "" && password != "" {

		_, err = app.DB.CreateUser(email, password)
		if err != nil {
			c.JSON(http.StatusBadRequest, app.errorJson("error signinu up"))
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, "email or password cannot be empty")
		return
	}

	signature, err := app.generateJwtToken(email)
	if err != nil {
		slog.Error("error generating token")
		return

	}

	c.JSON(http.StatusAccepted, app.successJson(signature))
}

func (app *application) loginPost(c *gin.Context) {
	//TODO : fix email and password validator
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, app.errorJson("error parsing form"))
		return
	}
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")
	if email == "" && password == "" {
		c.JSON(http.StatusBadRequest, app.errorJson("email or password cannot be empty"))
		return
	}
	_, err = app.DB.VerifyUser(email, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, app.errorJson("invalid criendientials"))
		slog.Info(err.Error())
		return
	}
	signature, err := app.generateJwtToken(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, app.errorJson("error generating token"))
	}
	c.JSON(http.StatusAccepted, app.successJson(signature))

}
