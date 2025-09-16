package handlers

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"githug.com/tejashwinn/spendsense-backend/internal/db/models"
)

type ExpenseHandler struct {
	DB *dynamodb.DynamoDB
}

func NewExpenseHandler() *ExpenseHandler {
	sess := session.Must(session.NewSession())
	db := dynamodb.New(sess, aws.NewConfig().WithRegion("us-west-2")) // Change region as needed
	return &ExpenseHandler{DB: db}
}

func (h *ExpenseHandler) AddExpense(c *gin.Context) {
	var expense models.Expense
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Logic to save expense to DynamoDB
	// ...

	c.JSON(http.StatusCreated, expense)
}

func (h *ExpenseHandler) GetGroupBalances(c *gin.Context) {
	groupID := c.Param("id")

	// Logic to retrieve group balances from DynamoDB
	// ...

	c.JSON(http.StatusOK, gin.H{"group_id": groupID, "balances": "..."}) // Replace with actual balances
}
