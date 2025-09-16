package models

import (
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
)

type Group struct {
	ID        string    `json:"id" dynamodbav:"id"`
	Name      string    `json:"name" dynamodbav:"name"`
	Currency  string    `json:"currency" dynamodbav:"currency"`
	CreatedBy string    `json:"created_by" dynamodbav:"created_by"`
	CreatedAt time.Time `json:"created_at" dynamodbav:"created_at"`
}

func NewGroup(name, createdBy string) *Group {
	return &Group{
		ID:        uuid.New().String(),
		Name:      name,
		Currency:  "USD",
		CreatedBy: createdBy,
		CreatedAt: time.Now(),
	}
}

func (g *Group) Save(db *dynamodb.DynamoDB) error {
	// Implementation for saving the group to DynamoDB
	return nil
}

func GetGroup(db *dynamodb.DynamoDB, id string) (*Group, error) {
	// Implementation for retrieving a group from DynamoDB
	return nil, nil
}
