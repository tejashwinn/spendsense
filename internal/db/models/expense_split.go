package models

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type ExpenseSplit struct {
	ExpenseID string  `json:"expense_id" dynamodbav:"expense_id"`
	UserID    string  `json:"user_id" dynamodbav:"user_id"`
	Amount    float64 `json:"amount" dynamodbav:"amount"`
}

func NewExpenseSplit(expenseID, userID string, amount float64) *ExpenseSplit {
	return &ExpenseSplit{
		ExpenseID: expenseID,
		UserID:    userID,
		Amount:    amount,
	}
}

func (es *ExpenseSplit) Save(db *dynamodb.DynamoDB) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("ExpenseSplits"),
		Item: map[string]*dynamodb.AttributeValue{
			"expense_id": {
				S: aws.String(es.ExpenseID),
			},
			"user_id": {
				S: aws.String(es.UserID),
			},
			"amount": {
				N: aws.String(fmt.Sprintf("%.2f", es.Amount)),
			},
		},
	}

	_, err := db.PutItem(input)
	return err
}
