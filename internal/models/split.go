package models

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
