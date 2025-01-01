package routes

import (
	"github.com/jinzhu/gorm"
	controllers "restaraunt_golang/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(routes *gin.Engine, db *gorm.DB) {
	routes.GET("/orderItems", controllers.GetOrderItems(db))
	routes.GET("/orderItems/:orderItem_id", controllers.GetOrderItem(db))
	routes.GET("/orderItems-order/:order_id", controllers.GetOrderItemsByOrder(db))
	routes.POST("/orderItems", controllers.CreateOrderItem(db))
	routes.PATCH("/orderItems/:orderItem_id", controllers.UpdateOrderItem(db))
}
