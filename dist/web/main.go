package main

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type application struct{
	router *gin.Engine
}

func main() {
	const PORT string = ":8000"
	app := &application{
		router : gin.Default(),
	}
	slog.Info(fmt.Sprintf("server staring on port %v", PORT))
	app.router.Run(PORT)
}
