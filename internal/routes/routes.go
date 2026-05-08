package routes

import (
	"auth-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine){
	router.GET("/health",handler.HealthCheck)
	router.POST("/user", handler.CreateUser)
} 