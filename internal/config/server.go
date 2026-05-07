package config

import (
	"github.com/gin-gonic/gin"
)


func SetUpRouter() *gin.Engine {
	//create default gin router

	router := gin.Default()
	return router
}