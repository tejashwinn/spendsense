package models

import (
	"time"
)

type Expense struct {
	ID          uint `gorm:"primaryKey"`
	GroupID     uint
	Group       Group `gorm:"foreignKey:GroupID"`
	CreatedBy   uint  `gorm:"not null"`
	User        User  `gorm:"foreignKey:CreatedBy"`
	Description string
	TotalAmount float64   `gorm:"not null"`
	Currency    string    `gorm:"not null"`
	Date        time.Time `gorm:"not null"`
	IsRecurring bool      `gorm:"default:false"`
	CreatedAt   time.Time
}
