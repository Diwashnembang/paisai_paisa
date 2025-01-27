package main

import "github.com/gin-gonic/gin"

func (app *application)home(c *gin.Context) {
	c.String(200, "hello world")
}
