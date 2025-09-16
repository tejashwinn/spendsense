package handlers

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"githug.com/tejashwinn/spendsense-backend/internal/db/models"
)

type GroupHandler struct {
	DB *dynamodb.DynamoDB
}

func NewGroupHandler() *GroupHandler {
	sess := session.Must(session.NewSession())
	db := dynamodb.New(sess, aws.NewConfig().WithRegion("us-west-2"))
	return &GroupHandler{DB: db}
}

func (h *GroupHandler) CreateGroup(c *gin.Context) {
	var group models.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Logic to save the group to DynamoDB
	// ...

	c.JSON(http.StatusCreated, group)
}

func (h *GroupHandler) AddMember(c *gin.Context) {
	groupID := c.Param("id")
	var member models.GroupMember
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Logic to add member to the group in DynamoDB
	// ...

	c.JSON(http.StatusOK, gin.H{"message": "Member added successfully"})
}
