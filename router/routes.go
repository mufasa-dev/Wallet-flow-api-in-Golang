package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mufasa-dev/Wallet-flow-api-in-Golang/handlers"
	"github.com/mufasa-dev/Wallet-flow-api-in-Golang/middleware"
)

func initializeRoutes(router *gin.Engine) {
	router.POST("/login", handlers.Login)

	v1 := router.Group("/api/v1")
	v1.Use(middleware.AuthMiddleware())
	{

		v1.GET("user", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Opening",
			})
		})
		v1.POST("user", handlers.CreateUserHandler)

		v1.PUT("user", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Opening",
			})
		})
		v1.POST("users", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Opening",
			})
		})
	}
}
