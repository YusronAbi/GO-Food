package controllers

import (
	"orderfoodonline/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	db *gorm.DB
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{db}
}

func (c *ProductController) GetAll(ctx *gin.Context) {
	var products []models.Product
	if err := c.db.Find(&products).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to fetch products"})
		return
	}

	ctx.JSON(200, products)
}

func (c *ProductController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	var product models.Product

	if err := c.db.First(&product, id).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	ctx.JSON(200, product)
}

func (c *ProductController) Create(ctx *gin.Context) {
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.db.Create(&product).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create product"})
		return
	}

	ctx.JSON(201, product)
}

func (c *ProductController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var product models.Product

	if err := c.db.First(&product, id).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.db.Save(&product).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to update product"})
		return
	}

	ctx.JSON(200, product)
}

func (c *ProductController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.db.Delete(&models.Product{}, id).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to delete product"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Product deleted successfully"})
}
