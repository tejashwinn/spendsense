package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response represents a standard API response structure
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// SendResponse sends a JSON response to the client
func SendResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Status:  http.StatusText(statusCode),
		Message: message,
		Data:    data,
	})
}

// SendError sends a JSON error response to the client
func SendError(c *gin.Context, statusCode int, message string) {
	SendResponse(c, statusCode, message, nil)
}