package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `gorm:"unique;not null;index"`
	PasswordHash string `gorm:"not null"`
	Transactions []Transaction
}

type Transaction struct {
	gorm.Model
	UserID      uint    `gorm:"not null"`
	Type        string  `gorm:"not null;check:type IN ('income', 'expense')"`
	Amount      float64 `gorm:"not null;type:decimal(10,2)"`
	Account     string  `gorm:"not null ;default:'cash'"`
	Description string
}
