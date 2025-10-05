package handlers

import (
	"net/http"

	"github.com/tejashwinn/spendsense/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateGroup godoc
// @Summary Create a new group
// @Description Create a new group
// @Tags groups
// @Accept json
// @Produce json
// @Param group body models.GroupRequest true "Group object"
// @Success 201 {object} models.GroupResponse
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
// @Success 200 {object} models.GroupResponse
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
// @Param group body models.GroupRequest true "Group object"
// @Success 200 {object} models.GroupResponse
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
// @Param id path int true "Group ID"
// @Param member body models.GroupMemberRequest true "GroupMember object"
// @Success 201 {object} models.GroupMemberResponse
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
