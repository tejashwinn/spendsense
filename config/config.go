package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds environment variables
type Config struct {
	DBUrl       string
	Port        string
	Environment string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return Config{
		DBUrl:       os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
		Environment: os.Getenv("ENVIRONMENT"),
	}
}
