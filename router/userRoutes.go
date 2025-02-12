package routes

import (
	"orderfoodonline/controllers"
	"orderfoodonline/middlewares"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userController *controllers.UserController
	router         *gin.Engine
}

func NewUserRoutes(router *gin.Engine, userController *controllers.UserController) *UserRoutes {
	return &UserRoutes{
		userController: userController,
		router:         router,
	}
}

func (r *UserRoutes) Setup() {
	userGroup := r.router.Group("/api/users")
	{
		userGroup.POST("/register", r.userController.Register)
		userGroup.POST("/login", r.userController.Login)

		authorized := userGroup.Use(middlewares.AuthMiddleware())
		{
			authorized.GET("/profile", r.userController.GetProfile)
			authorized.PUT("/profile", r.userController.UpdateProfile)
			authorized.GET("/orders", r.userController.GetUserOrders)
		}

		admin := userGroup.Use(middlewares.AuthMiddleware(), middlewares.AdminMiddleware())
		{
			admin.GET("", r.userController.GetAllUsers)
			admin.GET("/:id", r.userController.GetUserByID)
			admin.DELETE("/:id", r.userController.DeleteUser)
		}
	}
}
