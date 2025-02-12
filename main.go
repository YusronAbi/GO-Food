package main

import (
	"orderfoodonline/config"
	"orderfoodonline/database"
	"orderfoodonline/repository"
	"orderfoodonline/router"
	"orderfoodonline/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load konfigurasi dari environment
	cfg := config.NewEnvConfig()

	// Inisialisasi database
	db := database.Init(cfg, database.DBMigrator)

	// Inisialisasi repository
	authRepo := repository.NewAuthRepository(db)
	menuRepo := repository.NewMenuRepository(db)
	cartRepo := repository.NewCartRepository(db)
	orderRepo := repository.NewOrderRepository(db)

	// Inisialisasi service
	authService := service.NewAuthService(authRepo)
	menuService := service.NewMenuService(menuRepo) // Sesuaikan dengan nama yang benar
	cartService := service.NewCartService(cartRepo)
	orderService := service.NewOrderService(orderRepo)

	// Inisialisasi router
	r := gin.Default()
	router.SetupAuthRouter(r, authService)
	router.SetupUserRouter(r, db)
	router.SetupMenuRouter(r, db)
	router.SetupCartRouter(r, db, menuService, cartService)
	router.SetupOrderRouter(r, db, cartService)
	router.SetupReportRouter(r, db, orderService, orderRepo, cartRepo)

	// Endpoint root
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"App Name": "Shop App",
			"Author":   "Junx",
			"Version":  "1.0.0",
		})
	})

	// Jalankan server pada port 8080
	r.Run(":8080")
}
