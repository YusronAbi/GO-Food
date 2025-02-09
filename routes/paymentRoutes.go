package routes

import (
	"orderfoodonline/controllers"
	"orderfoodonline/middlewares"

	"github.com/gin-gonic/gin"
)

type PaymentRoutes struct {
	paymentController *controllers.PaymentController
	router            *gin.Engine
}

func NewPaymentRoutes(router *gin.Engine, paymentController *controllers.PaymentController) *PaymentRoutes {
	return &PaymentRoutes{
		paymentController: paymentController,
		router:            router,
	}
}

func (r *PaymentRoutes) Setup() {
	paymentGroup := r.router.Group("/api/payments")
	{
		// Protected routes
		authorized := paymentGroup.Use(middlewares.AuthMiddleware())
		{
			authorized.POST("", r.paymentController.ProcessPayment)
			authorized.GET("/my-payments", r.paymentController.GetPaymentHistory)
			authorized.GET("/:id", r.paymentController.GetPaymentDetail)
			authorized.POST("/:id/proof", r.paymentController.UploadPaymentProof)
		}

		// Admin routes
		admin := paymentGroup.Use(middlewares.AuthMiddleware(), middlewares.AdminMiddleware())
		{
			admin.GET("", r.paymentController.GetAllPayments)
			admin.PUT("/:id/verify", r.paymentController.VerifyPayment)
		}
	}
}
