package routes

import (
	"github.com/jinzhu/gorm"
	controllers "restaraunt_golang/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(routes *gin.Engine, DB *gorm.DB) {
	routes.GET("/foods", controllers.GetFoods(DB))
	routes.GET("/foods/:food_id", controllers.GetFood(DB))
	routes.POST("/foods", controllers.CreateFood(DB))
	routes.PATCH("/foods/:food_id", controllers.UpdateFood(DB))
}
