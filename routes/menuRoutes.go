package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	controllers "restaraunt_golang/controllers"
)

func MenuRoutes(routes *gin.Engine, db *gorm.DB) {
	routes.GET("/menus", controllers.GetMenus(db))
	routes.GET("/menus/:menu_id", controllers.GetMenu(db))
	routes.POST("/menus", controllers.CreateMenu(db))
	routes.PATCH("/menus/:menus_id", controllers.UpdateMenu(db))
}
