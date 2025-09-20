package handlers

import (
	"net/http"
	"spendsense/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateExpense godoc
// @Summary Create a new expense
// @Description Create a new expense
// @Tags expenses
// @Accept json
// @Produce json
// @Param expense body models.Expense true "Expense object"
// @Success 201 {object} models.Expense
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /expenses [post]
func CreateExpense(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var expense models.Expense
		if err := c.ShouldBindJSON(&expense); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&expense).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, expense)
	}
}

// GetExpense godoc
// @Summary Get expense by ID
// @Description Get details of an expense by ID
// @Tags expenses
// @Produce json
// @Param id path int true "Expense ID"
// @Success 200 {object} models.Expense
// @Failure 404 {object} map[string]string
// @Router /expenses/{id} [get]
func GetExpense(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var expense models.Expense
		id := c.Param("id")
		if err := db.Preload("Group").Preload("User").First(&expense, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "expense not found"})
			return
		}
		c.JSON(http.StatusOK, expense)
	}
}

// UpdateExpense godoc
// @Summary Update expense by ID
// @Description Update an existing expense
// @Tags expenses
// @Accept json
// @Produce json
// @Param id path int true "Expense ID"
// @Param expense body models.Expense true "Expense object"
// @Success 200 {object} models.Expense
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /expenses/{id} [put]
func UpdateExpense(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var expense models.Expense
		id := c.Param("id")
		if err := db.First(&expense, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "expense not found"})
			return
		}
		if err := c.ShouldBindJSON(&expense); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Save(&expense).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, expense)
	}
}

// DeleteExpense godoc
// @Summary Delete expense by ID
// @Description Delete an expense from the system
// @Tags expenses
// @Produce json
// @Param id path int true "Expense ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /expenses/{id} [delete]
func DeleteExpense(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Expense{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "expense deleted"})
	}
}
