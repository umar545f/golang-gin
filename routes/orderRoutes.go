package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	controllers "restaraunt_golang/controllers"
)

func OrderRoutes(routes *gin.Engine, db *gorm.DB) {
	routes.GET("/orders", controllers.GetOrders(db))
	routes.GET("/orders/:order_id", controllers.GetOrder(db))
	routes.POST("/orders", controllers.CreateOrder(db))
	routes.PATCH("/orders/:order_id", controllers.UpdateOrder(db))
}
