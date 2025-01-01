package routes

import (
	"github.com/jinzhu/gorm"
	controllers "restaraunt_golang/controllers"

	"github.com/gin-gonic/gin"
)

func TableRoutes(routes *gin.Engine, db *gorm.DB) {
	routes.GET("/tables", controllers.GetTables(db))
	routes.GET("/tables/:table_id", controllers.GetTable(db))
	routes.POST("/tables", controllers.CreateTable(db))
	routes.PATCH("/tables/:table_id", controllers.UpdateTable(db))
}
