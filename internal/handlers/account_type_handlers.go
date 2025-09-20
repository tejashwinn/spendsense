package handlers

import (
	"net/http"
	"spendsense/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AccountTypeHandler struct {
	DB *gorm.DB
}

// List all account types
func (h *AccountTypeHandler) ListAccountTypes(c *gin.Context) {
	var accountTypes []models.AccountType
	if err := h.DB.Find(&accountTypes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"account_types": models.AccountTypeToResponseList(accountTypes)})
}

// Get single account type by ID
func (h *AccountTypeHandler) GetAccountType(c *gin.Context) {
	id := c.Param("id")
	var accountType models.AccountType
	if err := h.DB.First(&accountType, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "account type not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, models.AccountTypeToResponse(&accountType))
}
