package util

type PageResponse[T any] struct {
	Items []T    `json:"items"`
	Total uint64 `json:"total"`
	Page  uint   `json:"page"`
	Size  uint   `json:"size"`
}
