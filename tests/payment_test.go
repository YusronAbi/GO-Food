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
	db.AutoMigrate(&models.Payment{})
	app = fiber.New()
	app.Post("/api/payments", createPayment)
}

func createPayment(c *fiber.Ctx) error {
	payment := new(models.Payment)
	if err := c.BodyParser(payment); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid input")
	}
	db.Create(&payment)
	return c.Status(http.StatusCreated).JSON(payment)
}

func TestCreatePayment(t *testing.T) {
	setup()

	newPayment := `{"orderId":1,"amount":19.98,"method":"credit_card"}`
	req := httptest.NewRequest("POST", "/api/payments", bytes.NewBufferString(newPayment))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("Failed to create payment: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", resp.StatusCode)
	}
}
