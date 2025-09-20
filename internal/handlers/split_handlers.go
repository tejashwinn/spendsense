package handlers

import (
	"net/http"
	"spendsense/internal/models"
	"spendsense/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddSplit godoc
// @Summary Add split to expense
// @Description Add a split to an expense
// @Tags splits
// @Accept json
// @Produce json
// @Param id path int true "Expense ID"
// @Param split body dto.SplitRequest true "Split object"
// @Success 201 {object} dto.SplitResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /expenses/{id}/splits [post]
func AddSplit(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var split models.Split
		id := c.Param("id")
		if err := c.ShouldBindJSON(&split); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		split.ExpenseID = util.ParseUint(id)
		if err := db.Create(&split).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, split)
	}
}

// GetSplits godoc
// @Summary Get splits for expense
// @Description Get all splits for an expense
// @Tags splits
// @Produce json
// @Param id path int true "Expense ID"
// @Success 200 {array} dto.SplitResponse
// @Failure 500 {object} map[string]string
// @Router /expenses/{id}/splits [get]
func GetSplits(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var splits []models.Split
		if err := db.Where("expense_id = ?", id).Find(&splits).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, splits)
	}
}
