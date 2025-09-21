package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// MonthlyReport godoc
// @Summary Get monthly report
// @Description Get monthly spend by category
// @Tags reports
// @Produce json
// @Success 200 {array} object
// @Failure 500 {object} map[string]string
// @Router /reports/monthly [get]
func MonthlyReport(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Example: monthly spend by category
		var result []struct {
			Month string
			Total float64
		}
		if err := db.Raw(`SELECT to_char(date, 'YYYY-MM') as month, SUM(total_amount) as total FROM expenses GROUP BY month ORDER BY month DESC`).Scan(&result).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

// TopSpenders godoc
// @Summary Get top spenders
// @Description Get top spenders by total amount
// @Tags reports
// @Produce json
// @Success 200 {array} object
// @Failure 500 {object} map[string]string
// @Router /reports/top-spenders [get]
func TopSpenders(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result []struct {
			UserID uint
			Total  float64
		}
		if err := db.Raw(`SELECT created_by as user_id, SUM(total_amount) as total FROM expenses GROUP BY created_by ORDER BY total DESC LIMIT 10`).Scan(&result).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
