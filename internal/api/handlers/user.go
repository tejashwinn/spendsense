package handlers

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"githug.com/tejashwinn/spendsense-backend/internal/db/models"
)

type UserHandler struct {
	DB *dynamodb.DynamoDB
}

func NewUserHandler() *UserHandler {
	sess := session.Must(session.NewSession())
	db := dynamodb.New(sess, aws.NewConfig().WithRegion("us-west-2")) // Change region as needed
	return &UserHandler{DB: db}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := user.Save(h.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := models.GetUserByID(h.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
