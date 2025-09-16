package models

import (
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Expense struct {
	ID          string    `json:"id" dynamodbav:"id"`
	GroupID     string    `json:"group_id" dynamodbav:"group_id"`
	Description string    `json:"description" dynamodbav:"description"`
	Amount      float64   `json:"amount" dynamodbav:"amount"`
	PaidBy      string    `json:"paid_by" dynamodbav:"paid_by"`
	Date        time.Time `json:"date" dynamodbav:"date"`
	Category    string    `json:"category" dynamodbav:"category"`
	CreatedAt   time.Time `json:"created_at" dynamodbav:"created_at"`
}

func (e *Expense) Save(db *dynamodb.DynamoDB) error {
	// Implementation for saving the expense to DynamoDB
	return nil
}

func GetExpenseByID(db *dynamodb.DynamoDB, id string) (*Expense, error) {
	// Implementation for retrieving an expense by ID from DynamoDB
	return nil, nil
}

func GetExpensesByGroupID(db *dynamodb.DynamoDB, groupID string) ([]Expense, error) {
	// Implementation for retrieving expenses by group ID from DynamoDB
	return nil, nil
}
