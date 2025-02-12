package router

import (
	"orderfoodonline/controller"
	"orderfoodonline/middleware"
	"orderfoodonline/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupMenuRouter mengatur route untuk menu
func SetupMenuRouter(r *gin.Engine, db *gorm.DB) {
	menuRepository := repository.NewMenuRepository(db)
	menuHandler := controller.NewMenuHandler(menuRepository)

	menuGroup := r.Group("/menus")
	menuGroup.Use(middleware.AuthProtected(db))
	{
		menuGroup.GET("", menuHandler.GetMany)
		menuGroup.Static("/image", "./uploads")
		menuGroup.GET("/download/:id", menuHandler.DownloadImage)
		menuGroup.GET("/:id", menuHandler.GetOne)
		menuGroup.POST("", middleware.RoleRequired("admin"), menuHandler.CreateOne)
		menuGroup.PUT("/:id", middleware.RoleRequired("admin"), menuHandler.UpdateOne)
		menuGroup.DELETE("/:id", middleware.RoleRequired("admin"), menuHandler.DeleteOne)
	}
}
