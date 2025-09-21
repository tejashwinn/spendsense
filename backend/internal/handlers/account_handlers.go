package handlers

import (
	"net/http"
	"time"

	"github.com/tejashwinn/spendsense/internal/models"
	"github.com/tejashwinn/spendsense/internal/repo"
	"github.com/tejashwinn/spendsense/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AccountHandler struct {
	AccountRepo  *repo.AccountRepo
	UserRepo     *repo.UserRepo
	CurrencyRepo *repo.CurrencyRepo
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

	userID := util.UserIDFromContext(c)
	accType, err := h.AccountRepo.GetAccountTypeByID(req.TypeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user, err := h.UserRepo.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	currency, err := h.CurrencyRepo.GetCurrencyByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "currency not found"})
		return
	}

	account := models.RequestToAccount(req, userID)
	account.Type = *accType
	account.User = *user
	account.Currency = *currency
	if err := h.AccountRepo.CreateAccount(&account); err != nil {
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
	account, err := h.AccountRepo.GetAccountByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, models.AccountToResponse(*account))
}

// List accounts
// @Summary List accounts
// @Description List all accounts for the user
// @Tags accounts
// @Produce json
// @Success 200 {object} util.AccountPageResponse
// @Router /accounts [get]
func (h *AccountHandler) ListAccounts(c *gin.Context) {
	userID := util.UserIDFromContext(c)
	accounts, err := h.AccountRepo.ListAccountsByUser(userID)
	if err != nil {
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

	account, err := h.AccountRepo.GetAccountByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	models.UpdateRequestToModel(account, req)
	account.UpdatedAt = time.Now()

	if err := h.AccountRepo.UpdateAccount(account); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.AccountToResponse(*account))
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
	if err := h.AccountRepo.DeleteAccount(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
