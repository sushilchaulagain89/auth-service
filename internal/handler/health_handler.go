package handler

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"auth-service/internal/service"
)

func HealthCheck(c *gin.Context){
	response := service.GetHealthStatus()
	c.JSON(http.StatusOK,response)
}