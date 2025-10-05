package handlers

import (
	"net/http"

	"github.com/tejashwinn/spendsense/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ActivityFeed godoc
// @Summary Get activity feed
// @Description Get recent expenses and settlements
// @Tags activity
// @Produce json
// @Success 200 {array} object
// @Failure 500 {object} map[string]string
// @Router /activity [get]
func ActivityFeed(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Example: recent expenses and settlements
		var activities []interface{}
		var expenses []models.Expense
		var settlements []models.Settlement
		db.Order("created_at desc").Limit(10).Find(&expenses)
		db.Order("created_at desc").Limit(10).Find(&settlements)
		for _, e := range expenses {
			activities = append(activities, e)
		}
		for _, s := range settlements {
			activities = append(activities, s)
		}
		c.JSON(http.StatusOK, activities)
	}
}
