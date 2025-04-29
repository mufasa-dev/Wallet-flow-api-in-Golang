package router

import (
	"net/http"

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

		v1.GET("user", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Opening",
			})
		})

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
