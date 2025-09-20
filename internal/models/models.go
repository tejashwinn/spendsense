package models

import (
	"time"

	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"uniqueIndex;not null"`
	Name      string
	CreatedAt time.Time
}

// Group model
type Group struct {
	gorm.Model

	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	OwnerID   uint   `gorm:"not null"`
	Owner     User   `gorm:"foreignKey:OwnerID"`
	CreatedAt time.Time
}

// GroupMember model
type GroupMember struct {
	gorm.Model

	ID       uint   `gorm:"primaryKey"`
	GroupID  uint   `gorm:"not null"`
	UserID   uint   `gorm:"not null"`
	Role     string `gorm:"default:member"`
	JoinedAt time.Time
	Group    Group `gorm:"foreignKey:GroupID"`
	User     User  `gorm:"foreignKey:UserID"`
}

// Expense model
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

// Split model
type Split struct {
	ID        uint    `gorm:"primaryKey"`
	ExpenseID uint    `gorm:"not null"`
	Expense   Expense `gorm:"foreignKey:ExpenseID"`
	UserID    uint    `gorm:"not null"`
	User      User    `gorm:"foreignKey:UserID"`
	SplitType string  `gorm:"not null"`
	Value     float64 `gorm:"not null"`
	Settled   bool    `gorm:"default:false"`
}

// Settlement model
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

// Comment model
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
