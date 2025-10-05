package models

import (
	"time"

	"gorm.io/gorm"
)

type AccountType struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(50);unique;not null"`
	Type      string    `gorm:"type:varchar(50);unique;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Response struct
type AccountTypeResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func AccountTypeToResponse(at *AccountType) *AccountTypeResponse {
	return &AccountTypeResponse{
		ID:   at.ID,
		Name: at.Name,
		Type: at.Type,
	}
}

func AccountTypeToResponseList(accountTypes []AccountType) []AccountTypeResponse {
	var responses []AccountTypeResponse
	for _, accountType := range accountTypes {
		responses = append(responses, *AccountTypeToResponse(&accountType))
	}
	return responses
}
