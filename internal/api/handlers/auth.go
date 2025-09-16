package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"githug.com/tejashwinn/spendsense-backend/internal/auth"
	"githug.com/tejashwinn/spendsense-backend/internal/db/models"
	"githug.com/tejashwinn/spendsense-backend/internal/utils"
)

type AuthHandler struct {
	UserModel models.User
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := h.UserModel.GetUserByEmail(nil, loginRequest.Email)
	if err != nil || !utils.CheckPasswordHash(loginRequest.Password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
