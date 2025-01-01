package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"restaraunt_golang/models"
)

func GetOrderItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var orderItems []models.OrderItem
		if err := db.Find(&orderItems).Error; err != nil {
			log.Println("Error fetching order items:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order items"})
			return
		}
		c.JSON(http.StatusOK, orderItems)
	}
}

func GetOrderItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderItemID := c.Param("orderItem_id")
		var orderItem models.OrderItem
		if err := db.Where("id = ?", orderItemID).First(&orderItem).Error; err != nil {
			log.Println("Error fetching order item:", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "Order item not found"})
			return
		}
		c.JSON(http.StatusOK, orderItem)
	}
}

func GetOrderItemsByOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderID := c.Param("order_id")
		var orderItems []models.OrderItem
		if err := db.Where("order_id = ?", orderID).Find(&orderItems).Error; err != nil {
			log.Println("Error fetching order items:", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "No items found for this order"})
			return
		}
		c.JSON(http.StatusOK, orderItems)
	}
}

func CreateOrderItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newOrderItem models.OrderItem
		if err := c.ShouldBindJSON(&newOrderItem); err != nil {
			log.Println("Error binding JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Create the order item
		if err := db.Create(&newOrderItem).Error; err != nil {
			log.Println("Error creating order item:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order item"})
			return
		}

		c.JSON(http.StatusCreated, newOrderItem)
	}
}

func UpdateOrderItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderItemID := c.Param("orderItem_id")
		var updatedOrderItem models.OrderItem
		if err := c.ShouldBindJSON(&updatedOrderItem); err != nil {
			log.Println("Error binding JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		var existingOrderItem models.OrderItem
		if err := db.Where("id = ?", orderItemID).First(&existingOrderItem).Error; err != nil {
			log.Println("Error fetching order item:", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "Order item not found"})
			return
		}

		// Update the order item
		if err := db.Model(&existingOrderItem).Updates(updatedOrderItem).Error; err != nil {
			log.Println("Error updating order item:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order item"})
			return
		}

		c.JSON(http.StatusOK, existingOrderItem)
	}
}
