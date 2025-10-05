package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID         uint        `gorm:"primaryKey" `
	Name       string      `gorm:"type:varchar(100);not null" `
	UserID     uint        `gorm:"not null;index" `
	User       User        `gorm:"foreignKey:UserID"`
	TypeID     uint        `gorm:"not null"`
	Type       AccountType `gorm:"foreignKey:TypeID"`
	Provider   string      `gorm:"type:varchar(100)"`
	Balance    float64     `gorm:"default:0"`
	CurrencyID uint        `gorm:"not null"`
	Currency   Currency    `gorm:"foreignKey:CurrencyID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Create account
type CreateAccountRequest struct {
	Name       string  `json:"name" binding:"required"`
	TypeID     uint    `json:"type_id" binding:"required"`
	Provider   string  `json:"provider,omitempty"`
	Balance    float64 `json:"balance,omitempty"`
	CurrencyID uint    `json:"currency,omitempty"`
}

// Update account
type UpdateAccountRequest struct {
	Name       string  `json:"name,omitempty"`
	TypeID     uint    `json:"type_id,omitempty"`
	Provider   string  `json:"provider,omitempty"`
	Balance    float64 `json:"balance,omitempty"`
	CurrencyID uint    `json:"currency,omitempty"`
}

// Single account response
type AccountResponse struct {
	ID        uint                `json:"id"`
	UserID    uint                `json:"user_id"`
	Name      string              `json:"name"`
	TypeID    uint                `json:"type_id"`
	Type      AccountTypeResponse `json:"type"`
	Provider  string              `json:"provider,omitempty"`
	Balance   float64             `json:"balance"`
	Currency  CurrencyResponse    `json:"currency"`
	CreatedAt string              `json:"created_at"`
	UpdatedAt string              `json:"updated_at"`
}

// List accounts response
type ListAccountsResponse struct {
	Accounts []AccountResponse `json:"accounts"`
}

// Map CreateAccountRequest -> Account
func RequestToAccount(req CreateAccountRequest, userID uint) Account {
	return Account{
		Name:       req.Name,
		TypeID:     req.TypeID,
		Provider:   req.Provider,
		Balance:    req.Balance,
		CurrencyID: req.CurrencyID,
		UserID:     userID,
	}
}

// Map UpdateAccountRequest -> update fields of Account
func UpdateRequestToModel(acc *Account, req UpdateAccountRequest) {
	if req.Name != "" {
		acc.Name = req.Name
	}
	if req.TypeID != 0 {
		acc.TypeID = req.TypeID
	}
	if req.Provider != "" {
		acc.Provider = req.Provider
	}
	if req.Balance != 0 {
		acc.Balance = req.Balance
	}
	if req.CurrencyID != 0 {
		acc.CurrencyID = req.CurrencyID
	}
}

// Map Account -> AccountResponse
func AccountToResponse(acc Account) AccountResponse {
	return AccountResponse{
		ID:        acc.ID,
		UserID:    acc.UserID,
		Name:      acc.Name,
		TypeID:    acc.TypeID,
		Type:      *AccountTypeToResponse(&acc.Type),
		Provider:  acc.Provider,
		Balance:   acc.Balance,
		Currency:  *CurrencyToResponse(&acc.Currency),
		CreatedAt: acc.CreatedAt.Format(time.RFC3339),
		UpdatedAt: acc.UpdatedAt.Format(time.RFC3339),
	}
}

// Map []Account -> ListAccountsResponse
func AccountsToListResponse(accounts []Account) []AccountResponse {
	accountResponses := make([]AccountResponse, len(accounts))
	for i, acc := range accounts {
		accountResponses[i] = AccountToResponse(acc)
	}
	return accountResponses
}
