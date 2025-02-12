package database

import (
	"fmt"
	"log"

	"orderfoodonline/entity"

	"gorm.io/gorm"
)

// DBMigrator menjalankan migrasi database
func DBMigrator(db *gorm.DB) error {
	entities := []any{
		&entity.User{},
		&entity.Menu{},
		&entity.Cart{},
		&entity.Order{},
	}

	for _, model := range entities {
		if err := db.AutoMigrate(model); err != nil {
			log.Printf("Migration failed for %T: %v\n", model, err)
			return fmt.Errorf("failed to migrate: %w", err)
		}
	}

	log.Println("[SUCCESS] Database migration completed successfully!")
	return nil
}
