package handler

import (
	"net/http"

	"auth-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserHandler struct {
	Service *service.UserService
}

func (h *UserHandler) CreateUser(c *gin.Context) {

	var req struct {
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil || req.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	err := h.Service.CreateUser(req.Email)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				c.JSON(409, gin.H{
					"error": "email already exists",
				})
				return
			}
		}
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}
