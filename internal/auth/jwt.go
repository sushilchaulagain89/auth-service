package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"errors"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID int
	Email  string
	jwt.RegisteredClaims
}

func GenerateAllTokens(uid int, email string) (accessToken string, refreshToken string, err error) {
	if len(jwtSecret) == 0 {
    return "", "", errors.New("JWT_SECRET is not set")
}
	now := time.Now()
	accessExp := now.Add(time.Hour * 24)
	refreshExp := now.Add(time.Hour * 24 * 7)
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

	accessToken, err = accessTokenObj.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = refreshTokenObj.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
