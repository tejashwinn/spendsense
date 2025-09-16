package db

import (
    "context"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
    DynamoDBClient *dynamodb.DynamoDB
)

func InitDynamoDB(region string) error {
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String(region),
    })
    if err != nil {
        return err
    }

    DynamoDBClient = dynamodb.New(sess)
    return nil
}

func CreateTable(ctx context.Context, tableName string, attributeDefinitions []*dynamodb.AttributeDefinition, keySchema []*dynamodb.KeySchemaElement, provisionedThroughput *dynamodb.ProvisionedThroughput) error {
    input := &dynamodb.CreateTableInput{
        TableName:            aws.String(tableName),
        AttributeDefinitions: attributeDefinitions,
        KeySchema:           keySchema,
        ProvisionedThroughput: provisionedThroughput,
    }

    _, err := DynamoDBClient.CreateTableWithContext(ctx, input)
    return err
}

func DeleteTable(ctx context.Context, tableName string) error {
    input := &dynamodb.DeleteTableInput{
        TableName: aws.String(tableName),
    }

    _, err := DynamoDBClient.DeleteTableWithContext(ctx, input)
    return err
}