package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"restaraunt_golang/models"
)

func GetMenus(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var menus []models.Menu
		if db == nil {
			log.Println("Database connection is not established.")
			return
		}
		err := db.Find(&menus).Error
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch menus"})
			return
		}

		if len(menus) == 0 {
			ctx.JSON(http.StatusOK, gin.H{"message": "No menus available"})
			return
		}

		ctx.JSON(http.StatusOK, menus)
	}
}

func GetMenu(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		menuId := ctx.Param("menu_id")

		var menu models.Menu
		if err := db.Where("id = ?", menuId).First(&menu).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch menu"})
			return
		}
		ctx.JSON(http.StatusOK, menu)
	}
}

// CreateMenu handles creating a new menu and associating foods.
func CreateMenu(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var menu models.Menu

		if err := ctx.ShouldBindJSON(&menu); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
			return
		}

		//fetch all existing foods
		var existingFood []models.Food
		if err := db.Find(&existingFood).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch foods from the database"})
			return
		}

		var filteredFood []models.Food

		//check if any existing food already exist that is being added in menu
		//If yes don't add it
		for _, food := range menu.Foods {
			for _, existingFood := range existingFood {
				if food.Name == existingFood.Name {
					existingFood.OnMenu = true
					if err := db.Save(&existingFood).Error; err != nil {
						ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update food record"})
						return
					}

				} else {
					if err := db.Create(&food).Error; err != nil {
						ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create new food record"})
						return
					}
					filteredFood = append(filteredFood, food)
				}
			}
		}

		menu.Foods = filteredFood

		if err := db.Create(&menu).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save menu in db"})
			return
		}

		ctx.JSON(http.StatusCreated, menu)
	}
}

func UpdateMenu(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var menu models.Menu
		var updatedMenu models.Menu
		menuId := ctx.Param("menu_id")

		if err := ctx.ShouldBindJSON(&updatedMenu); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
			return
		}

		if err := db.Where("id = ?", menuId).First(&menu).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Food not found"})
			return
		}

		if err := db.Model(&menu).Updates(&updatedMenu).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update menu in db"})
			return
		}
		ctx.JSON(http.StatusCreated, updatedMenu)
	}
}
