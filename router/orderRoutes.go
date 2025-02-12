package routes

import (
	"orderfoodonline/middlewares"

	"github.com/gin-gonic/gin"
)

type OrderRoutes struct {
	orderController *controllers.OrderController
	router          *gin.Engine
}

func NewOrderRoutes(router *gin.Engine, orderController *controllers.OrderController) *OrderRoutes {
	return &OrderRoutes{
		orderController: orderController,
		router:          router,
	}
}

func (r *OrderRoutes) Setup() {
	orderGroup := r.router.Group("/api/orders")
	{
		authorized := orderGroup.Use(middlewares.AuthMiddleware())
		{
			authorized.POST("", r.orderController.Create)
			authorized.GET("/my-orders", r.orderController.GetUserOrders)
			authorized.GET("/:id", r.orderController.GetOrderDetail)
			authorized.PUT("/:id/cancel", r.orderController.CancelOrder)
		}

		admin := orderGroup.Use(middlewares.AuthMiddleware(), middlewares.AdminMiddleware())
		{
			admin.GET("", r.orderController.GetAllOrders)
			admin.PUT("/:id/status", r.orderController.UpdateStatus)
		}
	}
}
