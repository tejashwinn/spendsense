package handlers

import (
	"fmt"
	"net/http"
	"spendsense/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// User Handlers
func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, user)
	}
}
func GetUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		id := c.Param("id")
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		id := c.Param("id")
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.User{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
	}
}

// Group Handlers
func CreateGroup(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var group models.Group
		if err := c.ShouldBindJSON(&group); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&group).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, group)
	}
}
func GetGroup(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var group models.Group
		id := c.Param("id")
		if err := db.Preload("Owner").First(&group, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "group not found"})
			return
		}
		c.JSON(http.StatusOK, group)
	}
}
func UpdateGroup(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var group models.Group
		id := c.Param("id")
		if err := db.First(&group, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "group not found"})
			return
		}
		if err := c.ShouldBindJSON(&group); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Save(&group).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, group)
	}
}
func DeleteGroup(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Group{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "group deleted"})
	}
}
func AddGroupMember(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var member models.GroupMember
		if err := c.ShouldBindJSON(&member); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&member).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, member)
	}
}
func RemoveGroupMember(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		groupID := c.Param("id")
		userID := c.Param("userId")
		if err := db.Where("group_id = ? AND user_id = ?", groupID, userID).Delete(&models.GroupMember{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "member removed"})
	}
}

// Expense Handlers
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

// Split Handlers
func AddSplit(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var split models.Split
		id := c.Param("id")
		if err := c.ShouldBindJSON(&split); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		split.ExpenseID = parseUint(id)
		if err := db.Create(&split).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, split)
	}
}
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

// Settlement Handlers
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

// Comment Handlers
func AddComment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var comment models.Comment
		id := c.Param("id")
		if err := c.ShouldBindJSON(&comment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		comment.ExpenseID = parseUint(id)
		if err := db.Create(&comment).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, comment)
	}
}
func GetComments(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var comments []models.Comment
		if err := db.Where("expense_id = ?", id).Find(&comments).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, comments)
	}
}

// Reports & Analytics
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

// Audit log & Activity Feed
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

// Helper

func parseUint(s string) uint {
	var v uint
	fmt.Sscanf(s, "%d", &v)
	return v
}
