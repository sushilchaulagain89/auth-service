package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func getJWTSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

// ✅ SINGLE SOURCE OF TRUTH for claims
type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// Generate both access + refresh tokens
func GenerateAllTokens(uid int, email string) (accessToken string, refreshToken string, err error) {

	now := time.Now()

	accessExp := now.Add(24 * time.Hour)
	refreshExp := now.Add(7 * 24 * time.Hour)

	accessClaims := &Claims{
		UserID: uid,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExp),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	refreshClaims := &Claims{
		UserID: uid,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExp),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	accessTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	accessToken, err = accessTokenObj.SignedString(getJWTSecret())
	if err != nil {
		return "", "", err
	}

	refreshToken, err = refreshTokenObj.SignedString(getJWTSecret())
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
