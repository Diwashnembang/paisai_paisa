package database

import (
	"diwashnembnag/pisai-paisa-backend/internal/models"
	"errors"
	"fmt"
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
		slog.Error("error creating user %s", result.Error)
		return -1, errors.New("error creating user")
	} else {
		return int(user.ID), nil
	}

}

func (db *Crud) VerifyUser(email, password string) (bool, error) {

	var user models.User
	result := db.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		slog.Error("error finding user %s", result.Error)
		return false, fmt.Errorf("error finding user")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return false, fmt.Errorf("INCORECT PASSWORD")
	}
	
	return true, nil
}
