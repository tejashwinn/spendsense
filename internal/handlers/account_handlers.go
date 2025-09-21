package handlers

import (
	"net/http"
	"spendsense/internal/models"
	"spendsense/util"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AccountHandler struct {
	DB *gorm.DB
}

// Create account
// @Summary Create account
// @Description Create a new account for the user
// @Tags accounts
// @Accept json
// @Produce json
// @Param account body models.CreateAccountRequest true "Account info"
// @Success 201 {object} models.AccountResponse
// @Router /accounts [post]
func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var req models.CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Normally you’d extract userID from auth, here hardcoded for demo
	userID := uint(1)

	account := models.RequestToAccount(req, userID)
	if err := h.DB.Create(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, models.AccountToResponse(account))
}

// Get single account
// @Summary Get account
// @Description Get details of a single account by ID
// @Tags accounts
// @Produce json
// @Param id path int true "Account ID"
// @Success 200 {object} models.AccountResponse
// @Router /accounts/{id} [get]
func (h *AccountHandler) GetAccount(c *gin.Context) {
	id := c.Param("id")
	var account models.Account
	if err := h.DB.First(&account, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, models.AccountToResponse(account))
}

// List accounts
// @Summary List accounts
// @Description List all accounts for the user
// @Tags accounts
// @Produce json
// @Success 200 {object} util.AccountPageResponse
// @Router /accounts [get]
func (h *AccountHandler) ListAccounts(c *gin.Context) {
	var accounts []models.Account
	if err := h.DB.Find(&accounts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, util.PageResponse[models.AccountResponse]{
		Items: models.AccountsToListResponse(accounts),
		Total: uint64(len(accounts)),
	})
}

// Update account
// @Summary Update account
// @Description Update an existing account by ID
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path int true "Account ID"
// @Param account body models.UpdateAccountRequest true "Account update info"
// @Success 200 {object} models.AccountResponse
// @Router /accounts/{id} [put]
func (h *AccountHandler) UpdateAccount(c *gin.Context) {
	id := c.Param("id")
	var req models.UpdateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var account models.Account
	if err := h.DB.First(&account, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	models.UpdateRequestToModel(&account, req)
	account.UpdatedAt = time.Now()

	if err := h.DB.Save(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.AccountToResponse(account))
}

// Delete account
// @Summary Delete account
// @Description Delete an account by ID
// @Tags accounts
// @Param id path int true "Account ID"
// @Success 204 {string} string "No Content"
// @Router /accounts/{id} [delete]
func (h *AccountHandler) DeleteAccount(c *gin.Context) {
	id := c.Param("id")
	if err := h.DB.Delete(&models.Account{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
