package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `json:"name" binding:"required"`
	Description string         `json:"description"`
	Price       float64        `json:"price" binding:"required,gt=0"`
	Stock       int            `json:"stock" binding:"required,gte=0"`
	Category    string         `json:"category"`
	ImageURL    string         `json:"image_url"`
	IsAvailable bool           `json:"is_available" gorm:"default:true"`
	Orders      []Order        `json:"orders,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type ProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"required,gte=0"`
	Category    string  `json:"category"`
	ImageURL    string  `json:"image_url"`
}
