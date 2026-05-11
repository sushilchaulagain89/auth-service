package routes

import (
	"auth-service/internal/handler"
	"auth-service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, userHandler *handler.UserHandler) {

	// --------------------
	// Public Routes
	// --------------------
	router.POST("/user", userHandler.CreateUser)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// --------------------
	// Protected Routes (JWT required)
	// --------------------
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())

	// protected.GET("/profile", userHandler.Profile)
}