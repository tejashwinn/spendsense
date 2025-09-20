package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID        uint    `gorm:"primaryKey" `
	Name      string  `gorm:"type:varchar(100);not null" `
	UserID    uint    `gorm:"not null;index" `
	User      User    `gorm:"foreignKey:UserID"`
	Type      string  `gorm:"type:varchar(50);not null" `
	Provider  string  `gorm:"type:varchar(100)"`
	Balance   float64 `gorm:"default:0"`
	Currency  string  `gorm:"type:varchar(10);default:'INR'" `
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Create account
type CreateAccountRequest struct {
	Name     string  `json:"name" binding:"required"`
	Type     string  `json:"type" binding:"required"`
	Provider string  `json:"provider,omitempty"`
	Balance  float64 `json:"balance,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

// Update account
type UpdateAccountRequest struct {
	Name     string  `json:"name,omitempty"`
	Type     string  `json:"type,omitempty"`
	Provider string  `json:"provider,omitempty"`
	Balance  float64 `json:"balance,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

// Single account response
type AccountResponse struct {
	ID        uint    `json:"id"`
	UserID    uint    `json:"user_id"`
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Provider  string  `json:"provider,omitempty"`
	Balance   float64 `json:"balance"`
	Currency  string  `json:"currency"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

// List accounts response
type ListAccountsResponse struct {
	Accounts []AccountResponse `json:"accounts"`
}

// Map CreateAccountRequest -> Account
func RequestToAccount(req CreateAccountRequest, userID uint) Account {
	return Account{
		Name:     req.Name,
		Type:     req.Type,
		Provider: req.Provider,
		Balance:  req.Balance,
		Currency: req.Currency,
		UserID:   userID,
	}
}

// Map UpdateAccountRequest -> update fields of Account
func UpdateRequestToModel(acc *Account, req UpdateAccountRequest) {
	if req.Name != "" {
		acc.Name = req.Name
	}
	if req.Type != "" {
		acc.Type = req.Type
	}
	if req.Provider != "" {
		acc.Provider = req.Provider
	}
	if req.Balance != 0 {
		acc.Balance = req.Balance
	}
	if req.Currency != "" {
		acc.Currency = req.Currency
	}
}

// Map Account -> AccountResponse
func AccountToResponse(acc Account) AccountResponse {
	return AccountResponse{
		ID:        acc.ID,
		UserID:    acc.UserID,
		Name:      acc.Name,
		Type:      acc.Type,
		Provider:  acc.Provider,
		Balance:   acc.Balance,
		Currency:  acc.Currency,
		CreatedAt: acc.CreatedAt.Format(time.RFC3339),
		UpdatedAt: acc.UpdatedAt.Format(time.RFC3339),
	}
}

// Map Account -> AccountResponse
func MapAccountToResponse(acc Account) AccountResponse {
	return AccountResponse{
		ID:        acc.ID,
		UserID:    acc.UserID,
		Name:      acc.Name,
		Type:      acc.Type,
		Provider:  acc.Provider,
		Balance:   acc.Balance,
		Currency:  acc.Currency,
		CreatedAt: acc.CreatedAt.Format(time.RFC3339),
		UpdatedAt: acc.UpdatedAt.Format(time.RFC3339),
	}
}

// Map []Account -> ListAccountsResponse
func AccountsToListResponse(accounts []Account) ListAccountsResponse {
	responses := make([]AccountResponse, len(accounts))
	for i, acc := range accounts {
		responses[i] = AccountToResponse(acc)
	}
	return ListAccountsResponse{Accounts: responses}
}
