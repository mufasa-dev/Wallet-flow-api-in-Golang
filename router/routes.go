package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mufasa-dev/Wallet-flow-api-in-Golang/handlers"
	"github.com/mufasa-dev/Wallet-flow-api-in-Golang/middleware"
)

func initializeRoutes(router *gin.Engine) {
	handlers.InitializeHandler()

	router.POST("/sigin", handlers.Login)
	router.POST("/sigup", handlers.CreateUserHandler)

	v1 := router.Group("/api/v1")
	v1.Use(middleware.AuthMiddleware())
	{
		v1.GET("users", handlers.ListUserHandler)
		v1.GET("user", handlers.ShowUserHandler)
		v1.PUT("user", handlers.UpdateUserHandler)
		v1.DELETE("user", handlers.DeleteUserHandler)
		v1.POST("withdraw", handlers.WithdrawHandler)
	}
}
