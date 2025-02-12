package helper

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// GetUserIDFromCookie mengambil user ID dari token yang tersimpan dalam cookie
func GetUserIDFromCookie(ctx *gin.Context) (uint, error) {
	cookie, err := ctx.Cookie("token")
	if err != nil {
		return 0, fmt.Errorf("authorization token is missing")
	}

	token, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("failed to parse token claims")
	}

	userIDFloat, ok := claims["id"].(float64)
	if !ok {
		return 0, fmt.Errorf("missing or invalid user ID in token")
	}

	return uint(userIDFloat), nil
}

// GetRoleFromToken mengambil role user dari token yang tersimpan dalam cookie
func GetRoleFromToken(ctx *gin.Context) (string, error) {
	cookie, err := ctx.Cookie("token")
	if err != nil {
		return "", fmt.Errorf("authorization token is missing")
	}

	token, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("failed to parse token claims")
	}

	role, ok := claims["role"].(string)
	if !ok || role == "" {
		return "", fmt.Errorf("missing or invalid role in token")
	}

	return role, nil
}
