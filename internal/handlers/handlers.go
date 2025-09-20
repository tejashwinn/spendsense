package handlers

import (
	"fmt"
	"net/http"
	"spendsense/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// User Handlers
// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user in the system
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
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

// GetUser godoc
// @Summary Get user by ID
// @Description Get details of a user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
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

// UpdateUser godoc
// @Summary Update user by ID
// @Description Update an existing user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User object"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [put]
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

// DeleteUser godoc
// @Summary Delete user by ID
// @Description Delete a user from the system
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [delete]
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
// CreateGroup godoc
// @Summary Create a new group
// @Description Create a new group
// @Tags groups
// @Accept json
// @Produce json
// @Param group body models.Group true "Group object"
// @Success 201 {object} models.Group
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /groups [post]
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

// GetGroup godoc
// @Summary Get group by ID
// @Description Get details of a group by ID
// @Tags groups
// @Produce json
// @Param id path int true "Group ID"
// @Success 200 {object} models.Group
// @Failure 404 {object} map[string]string
// @Router /groups/{id} [get]
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

// UpdateGroup godoc
// @Summary Update group by ID
// @Description Update an existing group
// @Tags groups
// @Accept json
// @Produce json
// @Param id path int true "Group ID"
// @Param group body models.Group true "Group object"
// @Success 200 {object} models.Group
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /groups/{id} [put]
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

// DeleteGroup godoc
// @Summary Delete group by ID
// @Description Delete a group from the system
// @Tags groups
// @Produce json
// @Param id path int true "Group ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /groups/{id} [delete]
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

// AddGroupMember godoc
// @Summary Add member to group
// @Description Add a user to a group
// @Tags groups
// @Accept json
// @Produce json
// @Param member body models.GroupMember true "GroupMember object"
// @Success 201 {object} models.GroupMember
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /groups/{id}/members [post]
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

// RemoveGroupMember godoc
// @Summary Remove member from group
// @Description Remove a user from a group
// @Tags groups
// @Produce json
// @Param id path int true "Group ID"
// @Param userId path int true "User ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /groups/{id}/members/{userId} [delete]
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

// Split Handlers
// AddSplit godoc
// @Summary Add split to expense
// @Description Add a split to an expense
// @Tags splits
// @Accept json
// @Produce json
// @Param id path int true "Expense ID"
// @Param split body models.Split true "Split object"
// @Success 201 {object} models.Split
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
		split.ExpenseID = parseUint(id)
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
// @Success 200 {array} models.Split
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

// Settlement Handlers
// CreateSettlement godoc
// @Summary Create a new settlement
// @Description Create a new settlement
// @Tags settlements
// @Accept json
// @Produce json
// @Param settlement body models.Settlement true "Settlement object"
// @Success 201 {object} models.Settlement
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
// @Success 200 {object} models.Settlement
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

// Comment Handlers
// AddComment godoc
// @Summary Add comment to expense
// @Description Add a comment to an expense
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "Expense ID"
// @Param comment body models.Comment true "Comment object"
// @Success 201 {object} models.Comment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /expenses/{id}/comments [post]
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

// GetComments godoc
// @Summary Get comments for expense
// @Description Get all comments for an expense
// @Tags comments
// @Produce json
// @Param id path int true "Expense ID"
// @Success 200 {array} models.Comment
// @Failure 500 {object} map[string]string
// @Router /expenses/{id}/comments [get]
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

// Audit log & Activity Feed
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

// Helper

func parseUint(s string) uint {
	var v uint
	fmt.Sscanf(s, "%d", &v)
	return v
}
