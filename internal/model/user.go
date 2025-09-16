package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint64 `gorm:"primary_key;unique_index;AUTO_INCREMENT"`
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateUserReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRes struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}
