package util

type PageResponse[T any] struct {
	Items []T    `json:"items"`
	Total uint64 `json:"total,omitempty"`
	Page  uint   `json:"page,omitempty"`
	Size  uint   `json:"size,omitempty"`
}
