package config

import (
	"os"
)

type Config struct {
	AWSRegion         string
	DynamoDBTable     string
	JWTSecret         string
	PreferredCurrency string
}

func LoadConfig() (*Config, error) {
	return &Config{
		AWSRegion:         getEnv("AWS_REGION", "us-west-2"),
		DynamoDBTable:     getEnv("DYNAMODB_TABLE", "SplitwiseClone"),
		JWTSecret:         getEnv("JWT_SECRET", "your_jwt_secret"),
		PreferredCurrency: getEnv("PREFERRED_CURRENCY", "USD"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
