package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/mufasa-dev/Wallet-flow-api-in-Golang/docs"
	"github.com/mufasa-dev/Wallet-flow-api-in-Golang/handlers"
	"github.com/mufasa-dev/Wallet-flow-api-in-Golang/middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handlers.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	router.POST("/sigin", handlers.Login)
	router.POST("/sigup", handlers.CreateUserHandler)

	v1 := router.Group(basePath)
	v1.Use(middleware.AuthMiddleware())
	{
		v1.GET("users", handlers.ListUserHandler)
		v1.GET("user", handlers.ShowUserHandler)
		v1.PUT("user", handlers.UpdateUserHandler)
		v1.DELETE("user", handlers.DeleteUserHandler)
		v1.POST("withdraw", handlers.WithdrawHandler)
		v1.POST("deposit", handlers.DepositHandler)
		v1.POST("transfer", handlers.TransferHandler)
		v1.GET("statement", handlers.StatementHandler)
	}
	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
