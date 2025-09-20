package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"uniqueIndex;not null"`
	Name      string
	CreatedAt time.Time
}

// User DTOs
type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func UserToResponse(user *User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func RequestToUser(req *UserRequest) User {
	return User{
		Name:  req.Name,
		Email: req.Email,
	}
}

func UpdateUserFromRequest(user *User, req *UserRequest) {
	user.Name = req.Name
	user.Email = req.Email
}
