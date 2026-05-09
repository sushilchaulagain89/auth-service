package handler

import (
	"net/http"

	"auth-service/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *service.UserService
}

func (h *UserHandler) CreateUser(c *gin.Context) {

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil || req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "All fields required",
		})
		return
	}

	err := h.Service.CreateUser(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created",
	})
}

func (h *UserHandler) ValidateUser(c *gin.Context) {

	var req struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil || req.Email == ""|| req.Password == ""{
		c.JSON(http.StatusBadRequest, gin.H{"error": "All field required!!"})
		return
	}


	err := h.Service.Login(req.Email,req.Password)
	if err != nil{
		c.JSON(http.StatusUnauthorized,gin.H{
			"error": "invalid email/password",
		})
		return 
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"login success",
	})
}