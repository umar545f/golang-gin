package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"restaraunt_golang/models"
)

func GetTables(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tables []models.Table
		if err := db.Find(&tables).Error; err != nil {
			log.Println("Error fetching tables:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tables"})
			return
		}
		c.JSON(http.StatusOK, tables)
	}
}

func GetTable(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tableID := c.Param("table_id")
		var table models.Table
		if err := db.Where("id = ?", tableID).First(&table).Error; err != nil {
			log.Println("Error fetching table:", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "Table not found"})
			return
		}
		c.JSON(http.StatusOK, table)
	}
}

func CreateTable(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newTable models.Table
		if err := c.ShouldBindJSON(&newTable); err != nil {
			log.Println("Error binding JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Create the table
		if err := db.Create(&newTable).Error; err != nil {
			log.Println("Error creating table:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create table"})
			return
		}

		c.JSON(http.StatusCreated, newTable)
	}
}

func UpdateTable(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tableID := c.Param("table_id")
		var updatedTable models.Table
		if err := c.ShouldBindJSON(&updatedTable); err != nil {
			log.Println("Error binding JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		var existingTable models.Table
		if err := db.Where("id = ?", tableID).First(&existingTable).Error; err != nil {
			log.Println("Error fetching table:", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "Table not found"})
			return
		}

		if err := db.Model(&existingTable).Updates(updatedTable).Error; err != nil {
			log.Println("Error updating table:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update table"})
			return
		}

		c.JSON(http.StatusOK, existingTable)
	}
}
