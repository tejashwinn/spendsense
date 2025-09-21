package main

import (
	"log"

	"github.com/tejashwinn/spendsense/config"
	"github.com/tejashwinn/spendsense/database"
	"github.com/tejashwinn/spendsense/internal/routes"

	_ "github.com/tejashwinn/spendsense/docs" // Swagger docs import

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
