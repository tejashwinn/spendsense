package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID        uint    `gorm:"primaryKey"`
	ExpenseID uint    `gorm:"not null"`
	Expense   Expense `gorm:"foreignKey:ExpenseID"`
	UserID    uint    `gorm:"not null"`
	User      User    `gorm:"foreignKey:UserID"`
	Body      string
	CreatedAt time.Time
}

type CommentRequest struct {
	ExpenseID uint      `json:"expense_id"`
	UserID    uint      `json:"user_id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentResponse struct {
	ID        uint         `json:"id"`
	ExpenseID uint         `json:"expense_id"`
	Expense   UserResponse `json:"expense"`
	UserID    uint         `json:"user_id"`
	User      UserResponse `json:"user"`
	Body      string       `json:"body"`
	CreatedAt time.Time    `json:"created_at"`
}
