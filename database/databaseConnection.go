package database

import (
	"log"
	"restaraunt_golang/models"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/jinzhu/gorm"
)

// DB is global DB connection pool
var DB *gorm.DB

func InitializeDB() (*gorm.DB, error) {
	// Use host.docker.internal for Docker to host machine connection (for Docker Desktop)
	dsn := "root:Faster@1@tcp/restaraunt?parseTime=true"

	DB, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Printf("Unable to connect to database : %v", err)
		return nil, err
	}

	// Create tables if they do not exist
	DB.AutoMigrate(&models.Food{}, &models.Menu{}, &models.Order{}, &models.OrderItem{}, &models.Invoice{}, &models.Table{})

	log.Println("Database connected and migrated successfully")

	return DB, nil
}

func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Printf("Unable to close database")
	}
	log.Printf("Database closed")
}
