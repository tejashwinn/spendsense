package models

import (
	"time"

	"gorm.io/gorm"
)

type Settlement struct {
	gorm.Model
	ID        uint    `gorm:"primaryKey"`
	FromUser  uint    `gorm:"not null"`
	ToUser    uint    `gorm:"not null"`
	Amount    float64 `gorm:"not null"`
	Currency  string  `gorm:"not null"`
	CreatedAt time.Time
	From      User `gorm:"foreignKey:FromUser"`
	To        User `gorm:"foreignKey:ToUser"`
}
