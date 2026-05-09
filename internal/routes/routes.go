package routes

import (
	"auth-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine,userHandler *handler.UserHandler){
	router.POST("/user",userHandler.CreateUser)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
} 