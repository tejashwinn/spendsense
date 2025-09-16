package logic

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
)

type BalanceCalculator struct {
	DB *dynamodb.DynamoDB
}

func NewBalanceCalculator() *BalanceCalculator {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"), // Change to your desired region
	})
	if err != nil {
		log.Fatalf("failed to create session: %v", err)
	}

	db := dynamodb.New(sess)

	return &BalanceCalculator{DB: db}
}

func (bc *BalanceCalculator) CalculateNetBalance(expenses []Expense, splits []ExpenseSplit) map[string]float64 {
	netBalance := make(map[string]float64)

	for _, expense := range expenses {
		for _, split := range splits {
			if split.ExpenseID == expense.ID {
				netBalance[split.UserID] += split.Amount
				netBalance[expense.PaidBy] -= split.Amount
			}
		}
	}

	return netBalance
}