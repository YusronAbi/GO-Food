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
	db.AutoMigrate(&models.User{})

	app = fiber.New()
	app.Post("/api/users", createUser)
}

func createUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid input")
	}
	db.Create(&user)
	return c.Status(http.StatusCreated).JSON(user)
}

func TestCreateUser(t *testing.T) {
	setup()
	newUser := `{"name":"Devi","email":"deviana@gmail.com","password":"halodevi","phone":"1234567890","address":"Malang"}`
	req := httptest.NewRequest("POST", "/api/users", bytes.NewBufferString(newUser))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", resp.StatusCode)
	}
}
