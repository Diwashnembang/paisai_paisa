package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Email        string `gorm:"unique;not null"`
    PasswordHash string `gorm:"not null"`
    Transactions []Transaction
}

type Category struct {
    gorm.Model
    Name      string `gorm:"not null"`
    Type      string `gorm:"not null;check:type IN ('income', 'expense')"`
    UserID    uint   // For custom categories
    Transactions []Transaction
}

type Transaction struct {
    gorm.Model
    UserID     uint    `gorm:"not null"`
    CategoryID uint    `gorm:"not null"`
    Amount     float64 `gorm:"not null;type:decimal(10,2)"`
    Description string
    Date       string `gorm:"type:date;not null"`
}
