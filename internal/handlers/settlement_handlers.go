package handlers

import (
	"net/http"
	"spendsense/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateSettlement godoc
// @Summary Create a new settlement
// @Description Create a new settlement
// @Tags settlements
// @Accept json
// @Produce json
// @Param settlement body models.SettlementRequest true "Settlement object"
// @Success 201 {object} models.SettlementResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /settlements [post]
func CreateSettlement(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var settlement models.Settlement
		if err := c.ShouldBindJSON(&settlement); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&settlement).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, settlement)
	}
}

// GetSettlement godoc
// @Summary Get settlement by ID
// @Description Get details of a settlement by ID
// @Tags settlements
// @Produce json
// @Param id path int true "Settlement ID"
// @Success 200 {object} models.SettlementResponse
// @Failure 404 {object} map[string]string
// @Router /settlements/{id} [get]
func GetSettlement(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var settlement models.Settlement
		id := c.Param("id")
		if err := db.First(&settlement, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "settlement not found"})
			return
		}
		c.JSON(http.StatusOK, settlement)
	}
}
