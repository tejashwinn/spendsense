package util

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ParseUint(s string) uint {
	var v uint
	fmt.Sscanf(s, "%d", &v)
	return v
}

func UserIDFromContext(c *gin.Context) uint {
	var v uint
	fmt.Sscanf(c.GetHeader("x-user-id"), "%d", &v)
	return v
}
