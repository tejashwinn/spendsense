package models

import (
	"time"

	"gorm.io/gorm"
)

type Settlement struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey"`
	FromUser  uint      `gorm:"not null"`
	ToUser    uint      `gorm:"not null"`
	Amount    float64   `gorm:"not null"`
	Currency  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	From      User      `gorm:"foreignKey:FromUser"`
	To        User      `gorm:"foreignKey:ToUser"`
}

type SettlementRequest struct {
	FromUser uint    `json:"from_user" binding:"required"`
	ToUser   uint    `json:"to_user" binding:"required"`
	Amount   float64 `json:"amount" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
}

type SettlementResponse struct {
	ID        uint         `json:"id"`
	FromUser  uint         `json:"from_user"`
	ToUser    uint         `json:"to_user"`
	Amount    float64      `json:"amount"`
	Currency  string       `json:"currency"`
	CreatedAt time.Time    `json:"created_at"`
	From      UserResponse `json:"from_user_details"`
	To        UserResponse `json:"to_user_details"`
}
