package handler

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"aurh-service/internal/service"
)

func HealthCheck(c *gin.Context){
	response := service.GetHealthStatus()
	c.JSON(http.StatusOK,response)
}