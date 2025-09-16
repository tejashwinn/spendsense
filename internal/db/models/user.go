package models

import (
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
)

type User struct {
	ID                string    `json:"id" dynamodbav:"id"`
	Name              string    `json:"name" dynamodbav:"name"`
	Email             string    `json:"email" dynamodbav:"email"`
	PasswordHash      string    `json:"password_hash" dynamodbav:"password_hash"`
	PreferredCurrency string    `json:"preferred_currency" dynamodbav:"preferred_currency"`
	CreatedAt         time.Time `json:"created_at" dynamodbav:"created_at"`
}

func NewUser(name, email, passwordHash string) *User {
	return &User{
		ID:                uuid.New().String(),
		Name:              name,
		Email:             email,
		PasswordHash:      passwordHash,
		PreferredCurrency: "USD",
		CreatedAt:         time.Now(),
	}
}

func (u *User) Save(db *dynamodb.DynamoDB) error {
	// Implementation for saving the user to DynamoDB
	return nil
}

func (u *User) GetUserByID(db *dynamodb.DynamoDB, id string) (*User, error) {
	// Implementation for retrieving a user by ID from DynamoDB
	return nil, nil
}

func (u *User) GetUserByEmail(db *dynamodb.DynamoDB, email string) (*User, error) {
	// Implementation for retrieving a user by email from DynamoDB
	return nil, nil
}
