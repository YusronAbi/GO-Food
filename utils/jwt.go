package util

import "github.com/golang-jwt/jwt"

func GenerateJWT(claims jwt.Claims, secret string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
}
