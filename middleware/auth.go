package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"orderfoodonline/entity"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

// extractToken mengambil token JWT dari header Authorization
func extractToken(ctx *gin.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("missing authorization header")
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return "", errors.New("invalid token format")
	}

	return tokenParts[1], nil
}

// extractUserID mengekstrak user ID dari claims token JWT
func extractUserID(claims jwt.MapClaims) (uint, error) {
	userIDFloat, ok := claims["id"].(float64)
	if !ok {
		return 0, errors.New("invalid user ID in token")
	}
	return uint(userIDFloat), nil
}

// AuthProtected middleware untuk memeriksa token JWT dalam header Authorization
func AuthProtected(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr, err := extractToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
			ctx.Abort()
			return
		}

		secret := []byte(os.Getenv("JWT_SECRET"))
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secret, nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": "Unauthorized - Invalid token",
			})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": "Unauthorized - Invalid claims",
			})
			ctx.Abort()
			return
		}

		userID, err := extractUserID(claims)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
			ctx.Abort()
			return
		}

		var user entity.User
		if err := db.First(&user, userID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": "Unauthorized - User not found",
			})
			ctx.Abort()
			return
		}

		ctx.Set("userId", userID)
		ctx.Next()
	}
}
