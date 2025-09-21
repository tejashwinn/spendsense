package models

import (
	"time"

	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey"`
	GroupID     uint      `gorm:"group_id;not null"`
	Group       Group     `gorm:"foreignKey:GroupID"`
	CreatedBy   uint      `gorm:"not null"`
	User        User      `gorm:"foreignKey:CreatedBy"`
	Description string    `gorm:"description"`
	TotalAmount float64   `gorm:"not null"`
	Currency    string    `gorm:"not null"`
	Date        time.Time `gorm:"not null"`
	IsRecurring bool      `gorm:"default:false"`
	CreatedAt   time.Time
}

type ExpenseResponse struct {
	ID          uint          `json:"id"`
	GroupID     uint          `json:"group_id"`
	Group       GroupResponse `json:"group"`
	CreatedBy   uint          `json:"created_by"`
	User        UserResponse  `json:"user"`
	Description string        `json:"description"`
	TotalAmount float64       `json:"total_amount"`
	Currency    string        `json:"currency"`
	Date        time.Time     `json:"date"`
	IsRecurring bool          `json:"is_recurring"`
	CreatedAt   time.Time     `json:"created_at"`
}

type ExpenseRequest struct {
	ID          uint      `json:"id"`
	GroupID     uint      `json:"group_id"`
	CreatedBy   uint      `json:"created_by"`
	Description string    `json:"description"`
	TotalAmount float64   `json:"total_amount"`
	Currency    string    `json:"currency"`
	Date        time.Time `json:"date"`
	IsRecurring bool      `json:"is_recurring"`
}
