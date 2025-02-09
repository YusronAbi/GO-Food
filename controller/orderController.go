package controllers

import (
	"orderfoodonline/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderController struct {
	db *gorm.DB
}

func NewOrderController(db *gorm.DB) *OrderController {
	return &OrderController{db}
}

func (c *OrderController) Create(ctx *gin.Context) {
	var order models.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// product details
	var product models.Product
	if err := c.db.First(&product, order.ProductID).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	//  total price
	order.TotalPrice = float64(order.Quantity) * product.Price

	userID, _ := ctx.Get("user_id")
	order.UserID = userID.(uint)

	if err := c.db.Create(&order).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create order"})
		return
	}

	ctx.JSON(201, order)
}

func (c *OrderController) GetUserOrders(ctx *gin.Context) {
	userID, _ := ctx.Get("user_id")
	var orders []models.Order

	if err := c.db.Where("user_id = ?", userID).Preload("Product").Find(&orders).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to fetch orders"})
		return
	}

	ctx.JSON(200, orders)
}

func (c *OrderController) UpdateStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	var input struct {
		Status string `json:"status" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.db.Model(&models.Order{}).Where("id = ?", id).Update("status", input.Status).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to update order status"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Order status updated successfully"})
}
