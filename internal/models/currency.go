package models

import (
	"time"

	"gorm.io/gorm"
)

type Currency struct {
	gorm.Model
	ID            uint   `gorm:"primaryKey"`
	Code          string `gorm:"size:3;unique;not null"`
	Name          string `gorm:"not null"`
	Symbol        string `gorm:"not null"`
	DecimalPlaces int    `gorm:"not null;default:2"`
	IsActive      bool   `gorm:"not null;default:true"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type CurrencyResponse struct {
	ID            uint   `json:"id"`
	Code          string `json:"code"`
	Name          string `json:"name"`
	Symbol        string `json:"symbol"`
	DecimalPlaces int    `json:"decimal_places"`
}

func CurrencyToResponse(c *Currency) *CurrencyResponse {
	return &CurrencyResponse{
		ID:            c.ID,
		Code:          c.Code,
		Name:          c.Name,
		Symbol:        c.Symbol,
		DecimalPlaces: c.DecimalPlaces,
	}
}

func CurrenciesToResponse(currencies []Currency) []CurrencyResponse {
	responses := make([]CurrencyResponse, len(currencies))
	for i, c := range currencies {
		responses[i] = *CurrencyToResponse(&c)
	}
	return responses
}
