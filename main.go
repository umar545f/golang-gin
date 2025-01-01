package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"restaraunt_golang/database"
	routesPackage "restaraunt_golang/routes"
)

func main() {
	// Default port fallback
	port := os.Getenv("PORT")
	if port == "" {
		port = "7200"
	}

	//establish database connection
	DB, err := database.InitializeDB()
	defer database.CloseDB()
	if err != nil {
		log.Println("Failed to establish DB connection")
	}
	// Create a new Gin router
	router := gin.New()
	router.Use(gin.Logger())

	routesPackage.FoodRoutes(router, DB)
	routesPackage.MenuRoutes(router, DB)
	routesPackage.TableRoutes(router, DB)
	routesPackage.OrderRoutes(router, DB)
	routesPackage.OrderItemRoutes(router, DB)
	routesPackage.InvoiceRoutes(router, DB)

	// Start the server
	router.Run(":" + port)
}
