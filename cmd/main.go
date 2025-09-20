package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func main() {
	// Load AWS config (uses ~/.aws/config or env vars)
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("localhost"),
		config.WithBaseEndpoint("http://localhost:8000/"),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "dummy", SecretAccessKey: "dummy", SessionToken: "dummy",
				Source: "Hard-coded credentials; values are irrelevant for local DynamoDB",
			},
		}))

	if err != nil {
		log.Fatalln(err)
	}

	// Create DynamoDB client
	svc := dynamodb.NewFromConfig(cfg)

	_, err = svc.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		TableName: aws.String("MyTable1"),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("ID"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("Name"),
				KeyType:       types.KeyTypeHash,
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		}},
	)

	if err != nil {
		log.Fatalln(err)
	}

	// Example: put item
	_, err = svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("MyTable1"),
		Item: map[string]types.AttributeValue{
			"ID":   &types.AttributeValueMemberS{Value: "123"},
			"Name": &types.AttributeValueMemberS{Value: "Alice"},
		},
	})
	if err != nil {
		log.Fatalln(err)
	}
	_, err = svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("MyTable1"),
		Item: map[string]types.AttributeValue{
			"ID":   &types.AttributeValueMemberS{Value: "234"},
			"Name": &types.AttributeValueMemberS{Value: "Alice"},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Example: get item
	out, err := svc.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("MyTable1"),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: "123"},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Item:", out.Item["ID"])
	fmt.Println("Item:", out.Item["Name"])

}
