package routes

import (
	"orderfoodonline/controllers"
	"orderfoodonline/middlewares"

	"github.com/gin-gonic/gin"
)

type ProductRoutes struct {
	productController *controllers.ProductController
	router            *gin.Engine
}

func NewProductRoutes(router *gin.Engine, productController *controllers.ProductController) *ProductRoutes {
	return &ProductRoutes{
		productController: productController,
		router:            router,
	}
}

func (r *ProductRoutes) Setup() {
	productGroup := r.router.Group("/api/products")
	{
		productGroup.GET("", r.productController.GetAll)
		productGroup.GET("/:id", r.productController.GetByID)
		productGroup.GET("/categories", r.productController.GetCategories)

		admin := productGroup.Use(middlewares.AuthMiddleware(), middlewares.AdminMiddleware())
		{
			admin.POST("", r.productController.Create)
			admin.PUT("/:id", r.productController.Update)
			admin.DELETE("/:id", r.productController.Delete)
			admin.PUT("/:id/stock", r.productController.UpdateStock)
		}
	}
}
