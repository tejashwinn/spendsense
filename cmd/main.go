package main

import (
	"log"
	"spendsense/config"
	"spendsense/internal/models"

	"spendsense/internal/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()
	if cfg.DBUrl == "" {
		log.Fatal("DATABASE_URL not set")
	}
	db, err := gorm.Open(postgres.Open(cfg.DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	// Auto migrate models
	db.AutoMigrate(
		&models.User{},
		&models.Group{},
		&models.GroupMember{},
		&models.Expense{},
		&models.Split{},
		&models.Settlement{},
		&models.Comment{},
	)

	r := gin.Default()

	// Register all API routes
	importRoutes := func() {
		// This is a hack to force import for side effects
	}
	importRoutes()

	// Register routes
	routes.RegisterRoutes(r, db)

	r.Run(":" + cfg.Port)
}
