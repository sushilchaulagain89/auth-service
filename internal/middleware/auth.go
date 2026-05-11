package middleware

import (
	"auth-service/internal/auth"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func getJWTSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 1. Read Authorization header
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			c.Abort()
			return
		}

		// 2. Validate Bearer format
		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization format"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// 3. Parse + validate token
		token, err := jwt.ParseWithClaims(
			tokenString,
			&auth.Claims{},
			func(token *jwt.Token) (interface{}, error) {
				return getJWTSecret(), nil
			},
		)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		// 4. Extract claims safely
		claims, ok := token.Claims.(*auth.Claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			c.Abort()
			return
		}

		// 5. Inject into request context
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)

		// 6. Continue request
		c.Next()
	}
}