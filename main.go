package main

import (
	"orderfoodonline/config"
	"orderfoodonline/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Create a new Fiber app
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	// Start the server
	app.Listen(":3000")
}
