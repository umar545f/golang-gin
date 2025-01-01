package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"restaraunt_golang/models"
)

func GetFoods(DB *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var foods []models.Food
		if DB == nil {
			log.Println("Database connection is not established.")
			return
		}
		err := DB.Find(&foods).Error
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch foods"})
			return
		}

		if len(foods) == 0 {
			ctx.JSON(http.StatusOK, gin.H{"message": "No foods available"})
			return
		}

		ctx.JSON(http.StatusOK, foods)
	}
}

func GetFood(DB *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		foodId := ctx.Param("food_id")

		var food models.Food

		err := DB.Where("id = ?", foodId).First(&food).Error
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Food not found"})
			return
		}
		ctx.JSON(http.StatusOK, food)
	}
}

func CreateFood(DB *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var food models.Food

		if err := ctx.ShouldBindJSON(&food); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
			return
		}

		if err := DB.Create(&food).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save food in db"})
			return
		}

		ctx.JSON(http.StatusCreated, food)
	}
}

func UpdateFood(DB *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var food models.Food
		var updatedFood models.Food
		foodId := ctx.Param("food_id")

		if err := ctx.ShouldBindJSON(&updatedFood); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
			return
		}

		if err := DB.Where("id = ?", foodId).First(&food).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Food not found"})
			return
		}

		if err := DB.Model(&food).Updates(&updatedFood).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update food in db"})
			return
		}
		ctx.JSON(http.StatusCreated, updatedFood)

	}
}

func GetFoodByName(db *gorm.DB, foodName string) (*models.Food, error) {
	var food models.Food
	err := db.Where("name = ?", foodName).First(&food).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil // Food not found
		}
		return nil, err
	}

	// Return the found food record
	return &food, nil
}
