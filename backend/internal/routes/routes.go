package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/tejashwinn/spendsense/config"
	"github.com/tejashwinn/spendsense/internal/handlers"
	"github.com/tejashwinn/spendsense/internal/repo"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB, config *config.Config) {
	// Swagger UI endpoint

	RegisterSwaggerRoutes(config, r)
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: false,
	}))
	api := r.Group("/api")

	// User CRUD
	registerUserRoutes(db, api)

	// Account CRUD
	registerAccountRoutes(db, api)

	// Account Type CRUD
	registerAccountTypeRoutes(db, api)

	// Account Type CRUD
	registerCurrencyRoutes(db, api)

	api.GET("/oops", handlers.OopsHandler())

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

func registerCurrencyRoutes(db *gorm.DB, api *gin.RouterGroup) {
	currencyHandler := &handlers.CurrencyHandler{
		CurrencyRepo: repo.NewCurrencyRepo(db),
	}
	currencies := api.Group("/currencies")
	{
		currencies.GET("", currencyHandler.ListCurrencies)
		currencies.GET("/:id", currencyHandler.GetCurrency)
	}
}

func registerAccountTypeRoutes(db *gorm.DB, api *gin.RouterGroup) {
	accountTypeHandler := &handlers.AccountTypeHandler{
		AccountTypeRepo: repo.NewAccountTypeRepo(db),
	}
	accountTypes := api.Group("/account-types")
	{
		accountTypes.GET("", accountTypeHandler.ListAccountTypes)
		accountTypes.GET("/:id", accountTypeHandler.GetAccountType)
	}
}

func registerAccountRoutes(db *gorm.DB, api *gin.RouterGroup) {
	accountHandler := &handlers.AccountHandler{
		AccountRepo:  repo.NewAccountRepo(db),
		CurrencyRepo: repo.NewCurrencyRepo(db),
		UserRepo:     repo.NewUserRepo(db),
	}
	accounts := api.Group("/accounts")
	{
		accounts.POST("", accountHandler.CreateAccount)
		accounts.GET("", accountHandler.ListAccounts)
		accounts.GET("/:id", accountHandler.GetAccount)
		accounts.PUT("/:id", accountHandler.UpdateAccount)
		accounts.DELETE("/:id", accountHandler.DeleteAccount)
	}
}

func registerUserRoutes(db *gorm.DB, api *gin.RouterGroup) {
	userHandler := &handlers.UserHandler{
		UserRepo: repo.NewUserRepo(db),
	}
	users := api.Group("/users")
	{
		users.POST("", userHandler.CreateUser)
		users.GET("/:id", userHandler.GetUser)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
	}
}
