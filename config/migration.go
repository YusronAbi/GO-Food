package config

import (
	"log"
	"orderfoodonline/models"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	log.Println("Running database migration...")

	err := db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Order{},
		&models.Payment{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database migration completed successfully")

	seedInitialData(db)
}

func seedInitialData(db *gorm.DB) {
	var adminCount int64
	db.Model(&models.User{}).Where("role = ?", "admin").Count(&adminCount)

	if adminCount == 0 {
		admin := models.User{
			Name:     "Admin",
			Email:    "admin@example.com",
			Password: "admin123",
			Role:     "admin",
		}

		if err := db.Create(&admin).Error; err != nil {
			log.Printf("Failed to create admin user: %v", err)
		}
	}

	seedProducts(db)
}

func seedProducts(db *gorm.DB) {
	var productCount int64
	db.Model(&models.Product{}).Count(&productCount)

	if productCount == 0 {
		products := []models.Product{
			{
				Name:        "Nasi Goreng",
				Description: "Indonesian fried rice with special spices",
				Price:       25000,
				Stock:       100,
				ImageURL:    "nasigoreng.jpg",
			},
			{
				Name:        "Mie Goreng",
				Description: "Indonesian fried noodles with vegetables",
				Price:       20000,
				Stock:       100,
				ImageURL:    "miegoreng.jpg",
			},
		}

		for _, product := range products {
			if err := db.Create(&product).Error; err != nil {
				log.Printf("Failed to seed product %s: %v", product.Name, err)
			}
		}
	}
}
