package models

import (
	"time"

	"gorm.io/gorm"
)

type AccountType struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(50);unique;not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Response struct
type AccountTypeResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func AccountTypeToResponse(at *AccountType) *AccountTypeResponse {
	return &AccountTypeResponse{
		ID:   at.ID,
		Name: at.Name,
	}
}

func AccountTypeToResponseList(accountTypes []AccountType) []AccountTypeResponse {
	var responses []AccountTypeResponse
	for _, accountType := range accountTypes {
		responses = append(responses, *AccountTypeToResponse(&accountType))
	}
	return responses
}
