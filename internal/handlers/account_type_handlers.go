package handlers

import (
	"net/http"
	"spendsense/internal/models"
	"spendsense/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AccountTypeHandler struct {
	DB *gorm.DB
}

// List all account types
// @Summary List account types
// @Description List all account types
// @Tags account-types
// @Produce json
// @Success 200 {object} util.AccountTypePageResponse
// @Router /account-types [get]
func (h *AccountTypeHandler) ListAccountTypes(c *gin.Context) {
	var accountTypes []models.AccountType
	if err := h.DB.Find(&accountTypes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	accountTyperespones := models.AccountTypeToResponseList(accountTypes)
	c.JSON(http.StatusOK, util.PageResponse[models.AccountTypeResponse]{
		Items: accountTyperespones,
		Total: uint64(len(accountTypes)),
	})
}

// Get single account type by ID
// @Summary Get account type
// @Description Get a single account type by ID
// @Tags account-types
// @Produce json
// @Param id path int true "Account Type ID"
// @Success 200 {object} models.AccountTypeResponse
// @Router /account-types/{id} [get]
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
