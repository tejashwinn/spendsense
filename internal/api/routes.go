package api

import (
	"github.com/gin-gonic/gin"
	"githug.com/tejashwinn/spendsense-backend/internal/api/handlers"
	// "github.com/tejashwinn/spendsense-backend/internal/api/handlers"
)

func SetupRoutes(router *gin.Engine) {
	userHandler := handlers.UserHandler{}
	groupHandler := handlers.GroupHandler{}
	expenseHandler := handlers.ExpenseHandler{}
	authHandler := handlers.AuthHandler{}

	router.POST("/users", userHandler.CreateUser)
	router.POST("/auth/login", authHandler.Login)
	router.GET("/users/:id", userHandler.GetUser)

	router.POST("/groups", groupHandler.CreateGroup)
	router.POST("/groups/:id/members", groupHandler.AddMember)
	router.POST("/groups/:id/expenses", expenseHandler.AddExpense)
	router.GET("/groups/:id/balances", expenseHandler.GetGroupBalances)
	router.POST("/groups/:id/settle", expenseHandler.SettleGroup)
}
