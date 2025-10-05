package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tejashwinn/spendsense/config"
)

func RegisterSwaggerRoutes(config *config.Config, r *gin.Engine) {
	if config.Environment != "production" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
