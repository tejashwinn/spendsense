package main

import (
	"log"
	"spendsense/config"

	"spendsense/database"
	"spendsense/internal/routes"

	_ "spendsense/docs" // Swagger docs import

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	db, err := database.New(&cfg)
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	r := gin.Default()
	// Register routes
	routes.RegisterRoutes(r, db.DB, &cfg)

	r.Run(":" + cfg.Port)
}
