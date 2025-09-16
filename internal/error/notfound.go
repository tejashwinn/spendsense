package error

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tejashwinn/sependsense/internal/model"
)

// NotFound creates a gin middleware for handling page not found.
func NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, &model.Error{
			Error:            http.StatusText(http.StatusNotFound),
			ErrorCode:        http.StatusNotFound,
			ErrorDescription: "page not found",
		})
	}
}
