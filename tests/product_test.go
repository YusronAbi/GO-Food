package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"orderfoodonline/models"
	"testing"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var app *fiber.App
var db *gorm.DB

func setup() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Product{})

	app = fiber.New()
	app.Post("/api/products", createProduct)
}

func createProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid input")
	}
	db.Create(&product)
	return c.Status(http.StatusCreated).JSON(product)
}

func TestCreateProduct(t *testing.T) {
	setup()
	newProduct := `{"name":"Pizza","description":"Delicious cheese pizza","price":9.99,"stock":100}`
	req := httptest.NewRequest("POST", "/api/products", bytes.NewBufferString(newProduct))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to create product: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", resp.StatusCode)
	}
}
