package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"restaraunt_golang/models"
)

func GetOrders(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var orders []models.Order
		if err := db.Preload("OrderItems").Find(&orders).Error; err != nil {
			log.Println("Error fetching orders:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
			return
		}
		c.JSON(http.StatusOK, orders)
	}
}

func GetOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderID := c.Param("order_id")
		var order models.Order
		if err := db.Preload("OrderItems").Where("id = ?", orderID).First(&order).Error; err != nil {
			log.Println("Error fetching order:", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}
		c.JSON(http.StatusOK, order)
	}
}

func CreateOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newOrder models.Order
		if err := c.ShouldBindJSON(&newOrder); err != nil {
			log.Println("Error binding JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Create the order
		if err := db.Create(&newOrder).Error; err != nil {
			log.Println("Error creating order:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
			return
		}

		c.JSON(http.StatusCreated, newOrder)
	}
}

func UpdateOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderID := c.Param("order_id")
		var updatedOrder models.Order
		if err := c.ShouldBindJSON(&updatedOrder); err != nil {
			log.Println("Error binding JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		var existingOrder models.Order
		if err := db.Where("id = ?", orderID).First(&existingOrder).Error; err != nil {
			log.Println("Error fetching order:", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}

		// Update the order
		if err := db.Model(&existingOrder).Updates(updatedOrder).Error; err != nil {
			log.Println("Error updating order:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
			return
		}

		c.JSON(http.StatusOK, existingOrder)
	}
}
