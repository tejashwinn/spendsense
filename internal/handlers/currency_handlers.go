package handlers

import (
	"net/http"

	"github.com/tejashwinn/spendsense/internal/models"
	"github.com/tejashwinn/spendsense/internal/repo"
	"github.com/tejashwinn/spendsense/util"

	"github.com/gin-gonic/gin"
)

type CurrencyHandler struct {
	CurrencyRepo *repo.CurrencyRepo
}

// List all currencies
// @Summary List currencies
// @Description List all currencies
// @Tags account-types
// @Produce json
// @Success 200 {object} util.CurrencyPageResponse
// @Router /account-types [get]
func (h *CurrencyHandler) ListCurrencies(c *gin.Context) {
	currencies, err := h.CurrencyRepo.ListCurrencies()
	if err != nil {
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
	currency, err := h.CurrencyRepo.GetCurrencyByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "currency not found"})
	}
	c.JSON(http.StatusOK, models.CurrencyToResponse(currency))
}
