package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tejashwinn/sependsense/config"
	"github.com/tejashwinn/sependsense/internal/database"
)

func New(
	db *database.GormDatabase,
	config config.Configuration,
) (*gin.Engine, error) {
	router := gin.Default()
	userHandler := UserApi{
		DB: db,
	}
	router.POST("/users", userHandler.CreateUser)
	router.GET("/users/:id", userHandler.GetById)
	return router, nil
}
