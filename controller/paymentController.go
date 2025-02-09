package controllers

import (
	"orderfoodonline/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PaymentController struct {
	db *gorm.DB
}

func NewPaymentController(db *gorm.DB) *PaymentController {
	return &PaymentController{db}
}

func (c *PaymentController) ProcessPayment(ctx *gin.Context) {
	var payment models.Payment
	if err := ctx.ShouldBindJSON(&payment); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//  transaction
	tx := c.db.Begin()

	if err := tx.Create(&payment).Error; err != nil {
		tx.Rollback()
		ctx.JSON(500, gin.H{"error": "Failed to process payment"})
		return
	}

	if err := tx.Model(&models.Order{}).Where("id = ?", payment.OrderID).
		Updates(map[string]interface{}{
			"payment_id": payment.ID,
			"status":     "paid",
		}).Error; err != nil {
		tx.Rollback()
		ctx.JSON(500, gin.H{"error": "Failed to update order"})
		return
	}

	// Commit transaction
	tx.Commit()

	ctx.JSON(200, gin.H{
		"message": "Payment processed successfully",
		"payment": payment,
	})
}

func (c *PaymentController) GetPaymentHistory(ctx *gin.Context) {
	userID, _ := ctx.Get("user_id")
	var payments []models.Payment

	if err := c.db.Joins("Order").Where("orders.user_id = ?", userID).Find(&payments).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to fetch payment history"})
		return
	}

	ctx.JSON(200, payments)
}
