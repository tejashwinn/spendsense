package util

import "github.com/tejashwinn/spendsense/internal/models"

type PageResponse[T any] struct {
	Items []T    `json:"items,omitempty"`
	Total uint64 `json:"total,omitempty"`
	Page  uint   `json:"page,omitempty"`
	Size  uint   `json:"size,omitempty"`
}

type AccountPageResponse PageResponse[models.AccountResponse]

type AccountTypePageResponse PageResponse[models.AccountTypeResponse]

type CurrencyPageResponse PageResponse[models.CurrencyResponse]
