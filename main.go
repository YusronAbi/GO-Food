package main

import (
	"log"
	"orderfoodonline/config"
	"orderfoodonline/routes"
)

func main() {
	cfg := config.LoadConfig()

	db := config.InitDB(cfg)

	// Run Migration
	config.RunMigration(db)

	// Initialize Router
	r := routes.SetupRouter(db)

	log.Printf("Server is running on port %s", cfg.ServerPort)
	r.Run(":" + cfg.ServerPort)
}
