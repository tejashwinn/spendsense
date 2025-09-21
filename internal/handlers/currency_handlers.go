package handlers

import (
	"net/http"
	"spendsense/internal/models"
	"spendsense/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CurrencyHandler struct {
	DB *gorm.DB
}

// List all currencies
// @Summary List currencies
// @Description List all currencies
// @Tags account-types
// @Produce json
// @Success 200 {object} util.CurrencyPageResponse
// @Router /account-types [get]
func (h *CurrencyHandler) ListCurrencies(c *gin.Context) {
	var currencies []models.Currency
	if err := h.DB.Find(&currencies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	currencyResponses := models.CurrenciesToResponse(currencies)
	c.JSON(http.StatusOK, util.PageResponse[models.CurrencyResponse]{
		Items: currencyResponses,
		Total: uint64(len(currencies)),
	})
}

// Get single currency by ID
// @Summary Get Currency
// @Description Get a single Currency by ID
// @Tags account-types
// @Produce json
// @Param id path int true "Currency ID"
// @Success 200 {object} models.CurrencyResponse
// @Router /account-types/{id} [get]
func (h *CurrencyHandler) GetCurrency(c *gin.Context) {
	id := c.Param("id")
	var currency models.Currency
	if err := h.DB.First(&currency, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "account type not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, models.CurrencyToResponse(&currency))
}
