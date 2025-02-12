package entity

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	OrderID       uint           `json:"order_id"`
	Order         Order          `json:"order,omitempty" gorm:"foreignKey:OrderID"`
	Amount        float64        `json:"amount" binding:"required,gt=0"`
	PaymentMethod string         `json:"payment_method" binding:"required"`
	Status        string         `json:"status" gorm:"default:'pending'"`
	TransactionID string         `json:"transaction_id"`
	PaymentProof  string         `json:"payment_proof"`
	PaidAt        *time.Time     `json:"paid_at"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type PaymentRequest struct {
	OrderID       uint    `json:"order_id" binding:"required"`
	Amount        float64 `json:"amount" binding:"required,gt=0"`
	PaymentMethod string  `json:"payment_method" binding:"required"`
	PaymentProof  string  `json:"payment_proof"`
}

type PaymentVerification struct {
	PaymentID     uint   `json:"payment_id" binding:"required"`
	TransactionID string `json:"transaction_id" binding:"required"`
	Status        string `json:"status" binding:"required"`
}
