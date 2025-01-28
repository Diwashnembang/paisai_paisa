package database

import (
	"diwashnembnag/pisai-paisa-backend/internal/models"
	"errors"
	"log/slog"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Crud struct {
	DB *gorm.DB
}

func (db *Crud) CreateUser(email, password string) (int, error) {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}
	user := &models.User{
		Email:        email,
		PasswordHash: string(passwordHash),
	}
	result := db.DB.Create(user)
	if result.Error != nil {
		slog.Error("error creating user", result.Error)
		return -1, errors.New("error creating user")
	} else {
		return int(user.ID), nil
	}

}
