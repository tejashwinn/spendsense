package routes

import (
	"spendsense/internal/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	// Swagger UI endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")

	// User CRUD
	userHandler := &handlers.UserHandler{DB: db}
	users := api.Group("/users")
	{
		users.POST("/users", userHandler.CreateUser)
		users.GET("/users/:id", userHandler.GetUser)
		users.PUT("/users/:id", userHandler.UpdateUser)
		users.DELETE("/users/:id", userHandler.DeleteUser)
	}

	// Account CRUD
	accountHandler := &handlers.AccountHandler{DB: db}
	accounts := api.Group("/accounts")
	{
		accounts.POST("", accountHandler.CreateAccount)
		accounts.GET("", accountHandler.ListAccounts)
		accounts.GET("/:id", accountHandler.GetAccount)
		accounts.PUT("/:id", accountHandler.UpdateAccount)
		accounts.DELETE("/:id", accountHandler.DeleteAccount)
	}

	// Account CRUD
	accountTypeHandler := &handlers.AccountTypeHandler{DB: db}
	accountTypes := api.Group("/account-types")
	{
		accountTypes.GET("", accountTypeHandler.ListAccountTypes)
		accountTypes.GET("/:id", accountTypeHandler.GetAccountType)
	}

	// Group CRUD & membership
	api.POST("/groups", handlers.CreateGroup(db))
	api.GET("/groups/:id", handlers.GetGroup(db))
	api.PUT("/groups/:id", handlers.UpdateGroup(db))
	api.DELETE("/groups/:id", handlers.DeleteGroup(db))
	api.POST("/groups/:id/members", handlers.AddGroupMember(db))
	api.DELETE("/groups/:id/members/:userId", handlers.RemoveGroupMember(db))

	// Expenses CRUD
	api.POST("/expenses", handlers.CreateExpense(db))
	api.GET("/expenses/:id", handlers.GetExpense(db))
	api.PUT("/expenses/:id", handlers.UpdateExpense(db))
	api.DELETE("/expenses/:id", handlers.DeleteExpense(db))

	// Splits
	api.POST("/expenses/:id/splits", handlers.AddSplit(db))
	api.GET("/expenses/:id/splits", handlers.GetSplits(db))

	// Settlements
	api.POST("/settlements", handlers.CreateSettlement(db))
	api.GET("/settlements/:id", handlers.GetSettlement(db))

	// Comments
	api.POST("/expenses/:id/comments", handlers.AddComment(db))
	api.GET("/expenses/:id/comments", handlers.GetComments(db))

	// Reports & analytics
	api.GET("/reports/monthly", handlers.MonthlyReport(db))
	api.GET("/reports/topspenders", handlers.TopSpenders(db))

	// Audit log & activity feed
	api.GET("/activity", handlers.ActivityFeed(db))
	// ...existing code...
}
