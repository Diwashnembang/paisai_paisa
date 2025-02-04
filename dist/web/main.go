package main

import (
	"diwashnembnag/pisai-paisa-backend/internal/database"
	"diwashnembnag/pisai-paisa-backend/internal/models"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type application struct {
	router     *gin.Engine
	DB         *database.Crud
	JWT_SECRET string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env files: %v", err)
		os.Exit(1)
	}
	dsn := "host=localhost user=admin password=admin dbname=paisai_paisa"
	const PORT string = ":8000"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
		os.Exit(1)
	}
	err = db.AutoMigrate(&models.User{}, &models.Transaction{})
	if err != nil {
		log.Fatal("failed to migrate schema", err)
		os.Exit(1)
	}

	app := &application{
		router: gin.Default(),
		DB: &database.Crud{
			DB: db,
		},
		JWT_SECRET: os.Getenv("JWT_SECRET"),
	}
	app.routes()
	slog.Info(fmt.Sprintf("server staring on port %v", PORT))
	app.router.Run(PORT)
}
