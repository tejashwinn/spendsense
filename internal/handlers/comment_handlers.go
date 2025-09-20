package handlers

import (
	"net/http"
	"spendsense/internal/models"
	"spendsense/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddComment godoc
// @Summary Add comment to expense
// @Description Add a comment to an expense
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "Expense ID"
// @Param comment body dto.CommentRequest true "Comment object"
// @Success 201 {object} dto.CommentResponse
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
		comment.ExpenseID = util.ParseUint(id)
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
// @Success 200 {array} dto.CommentResponse
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
