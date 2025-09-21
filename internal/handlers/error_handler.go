package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OopsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Something went wrong! Please try again later."},
		)
	}
}
