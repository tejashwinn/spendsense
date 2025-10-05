package models

import (
	"time"

	"gorm.io/gorm"
)

type Split struct {
	gorm.Model
	ID        uint    `gorm:"primaryKey"`
	ExpenseID uint    `gorm:"not null"`
	Expense   Expense `gorm:"foreignKey:ExpenseID"`
	UserID    uint    `gorm:"not null"`
	User      User    `gorm:"foreignKey:UserID"`
	SplitType string  `gorm:"not null"`
	Value     float64 `gorm:"not null"`
	Settled   bool    `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SplitResponse struct {
	ID        uint            `json:"id"`
	ExpenseID uint            `json:"expense_id"`
	Expense   ExpenseResponse `json:"expense"`
	UserID    uint            `json:"user_id"`
	User      UserResponse    `json:"user"`
	SplitType string          `json:"split_type"`
	Value     float64         `json:"value"`
	Settled   bool            `json:"settled"`
}

type SplitRequest struct {
	ID        uint    `json:"id"`
	ExpenseID uint    `json:"expense_id"`
	UserID    uint    `json:"user_id"`
	SplitType string  `json:"split_type"`
	Value     float64 `json:"value"`
	Settled   bool    `json:"settled"`
}
