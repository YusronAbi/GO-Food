package controller

import (
	"net/http"
	"orderfoodonline/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// AuthController menangani request autentikasi
type AuthController struct {
	AuthService service.AuthService
}

var validate = validator.New()

// NewAuthController membuat instance AuthController
func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

// Login menangani login pengguna
func (c *AuthController) Login(ctx *gin.Context) {
	var req dto.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := validate.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	token, err := c.AuthService.Login(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("token", token, 3600*24, "/", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}

// Register menangani registrasi pengguna
func (c *AuthController) Register(ctx *gin.Context) {
	var req dto.RegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := validate.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	token, err := c.AuthService.Register(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("token", token, 3600*24, "/", "", false, true)
	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "token": token})
}

// Logout menghapus token pengguna
func (c *AuthController) Logout(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
