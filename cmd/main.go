package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/routes"
	"githug.com/tejashwinn/spendsense-backend/internal/config"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()

	if err != nil {
		panic("Failed to load config: " + err.Error())
	}

	// Initialize Gin router
	router := gin.Default()

	// Set up routes
	routes.SetupRoutes(router)

	// Start the server
	router.Run(":" + cfg.ServerPort)
}
