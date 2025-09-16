package database

import (
	"context"
	"log"

	"github.com/tejashwinn/sependsense/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type GormDatabase struct {
	DB *gorm.DB
}

type DynamoDatabase struct {
	DB *dynamodb.Client
}

func New(connection string) (*GormDatabase, error) {
	db, err := gorm.Open(postgres.Open(connection))
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&model.User{})
	return &GormDatabase{
		DB: db,
	}, nil

}

func NewDynamo(connection string) (*DynamoDatabase, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		// CHANGE THIS TO us-east-1 TO USE AWS proper
		config.WithRegion("localhost"),
		// Comment the below out if not using localhost
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:8000", SigningRegion: "localhost"}, nil // The SigningRegion key was what's was missing! D'oh.
			})),
	)
	if err != nil {
		return nil, err
	}
	print("temp")
	client := dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.Credentials = credentials.NewStaticCredentialsProvider("DUMMYIDEXAMPLE", "DUMMYIDEXAMPLE", "DUMMYIDEXAMPLE")
	})
	print("temp")

	// Example: List tables
	out, err := client.ListTables(context.TODO(), &dynamodb.ListTablesInput{})
	print("temp")
	if err != nil {
		log.Fatalf("failed to list tables, %v", err)
	}
	print("temp")
	println(out)

	return &DynamoDatabase{
		DB: client,
	}, nil

}
