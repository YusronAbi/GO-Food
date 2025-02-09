package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserID      uint           `json:"user_id"`
	User        User           `json:"user,omitempty" gorm:"foreignKey:UserID"`
	OrderItems  []OrderItem    `json:"order_items"`
	TotalAmount float64        `json:"total_amount"`
	Status      string         `json:"status" gorm:"default:'pending'"`
	PaymentID   *uint          `json:"payment_id"`
	Payment     *Payment       `json:"payment,omitempty" gorm:"foreignKey:PaymentID"`
	Address     string         `json:"delivery_address"`
	Notes       string         `json:"notes"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type OrderItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity" binding:"required,gt=0"`
	Price     float64 `json:"price"`
	Subtotal  float64 `json:"subtotal"`
}

type OrderRequest struct {
	Items   []OrderItemRequest `json:"items" binding:"required,dive"`
	Address string             `json:"delivery_address" binding:"required"`
	Notes   string             `json:"notes"`
}

type OrderItemRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,gt=0"`
}
