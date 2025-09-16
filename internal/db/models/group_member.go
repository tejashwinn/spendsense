package models

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type GroupMember struct {
	GroupID string `json:"group_id" dynamodbav:"group_id"`
	UserID  string `json:"user_id" dynamodbav:"user_id"`
	Role    string `json:"role" dynamodbav:"role"`
}

func (gm *GroupMember) Save(db *dynamodb.DynamoDB) error {
	// Implementation for saving GroupMember to DynamoDB
	return nil
}

func GetGroupMembers(db *dynamodb.DynamoDB, groupID string) ([]GroupMember, error) {
	// Implementation for retrieving group members from DynamoDB
	return nil, nil
}

func DeleteGroupMember(db *dynamodb.DynamoDB, groupID, userID string) error {
	// Implementation for deleting a group member from DynamoDB
	return nil, nil
}
