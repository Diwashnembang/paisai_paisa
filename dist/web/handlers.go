package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *application) home(c *gin.Context) {
	c.String(200, "hello world")
}

func (app *application) signUpPost(c *gin.Context) {
	//TODO : make third party sign in and up
	var userID int
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, app.errorJson(fmt.Sprintf("Failed to parse form : %s", err)))
		return
	}
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")
	if email != "" && password != "" {

		userID, err = app.DB.CreateUser(email, password)
		if err != nil {
			c.JSON(http.StatusBadRequest, app.errorJson("error signinu up"))
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, "email or password cannot be empty")
		return
	}

	signature, err := app.generateJwtToken(userID)
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
		fmt.Print("thsi is ", email, password)
		c.JSON(http.StatusBadRequest, app.errorJson("email or password cannot be empty"))
		return
	}
	_, userID, err := app.DB.VerifyUser(email, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, app.errorJson("invalid criendientials"))
		slog.Info(err.Error())
		return
	}
	signature, err := app.generateJwtToken(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, app.errorJson("error generating token"))
	}
	c.JSON(http.StatusAccepted, app.successJson(signature))

}

type Transaction struct {
	userID      uint
	category    string
	amount      float64
	account     string
	description string
}

func (app *application) createTransactionPost(c *gin.Context) {

	transaction := &Transaction{}
	if v, exists := c.Get("userID"); exists {
		if value, ok := v.(uint); ok {
			transaction.userID = value
		} else {
			slog.Error("error getting userID form the context: invalid userID type")
			fmt.Printf("this the type of userif %T \n", v)
			c.JSON(http.StatusInternalServerError, app.errorJson("error creating transaction"))
			return
		}
	} else {
		slog.Error("error getting userID form the context")
		c.JSON(http.StatusInternalServerError, app.errorJson("error creating transaction"))
		return
	}

	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, app.errorJson("error parsing form"))
		return
	}
	transaction.category = c.Request.FormValue("type")
	transaction.amount, err = strconv.ParseFloat(c.Request.FormValue("amount"), 64)
	if err != nil {
		slog.Error("error parsing amount")
		c.JSON(http.StatusInternalServerError, app.errorJson("error creating transaction"))
		return

	}
	transaction.description = c.Request.FormValue("description")
	transaction.account = c.Request.FormValue("account")

	_, err = app.DB.CreateTransaction(transaction.userID, transaction.amount, transaction.category, transaction.account, transaction.description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.errorJson("error creating transaction"))
		return
	}
	c.JSON(http.StatusAccepted, app.successJson(""))

}
