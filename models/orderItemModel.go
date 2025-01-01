package models

import "github.com/jinzhu/gorm"

type OrderItem struct {
	gorm.Model
	OrderID  uint    `json:"order_id" gorm:"not null"`
	FoodID   uint    `json:"food_id" gorm:"not null"`
	Quantity int32   `json:"quantity" gorm:"not null"`
	Price    float64 `json:"price" gorm:"not null"` // Price at the time of the order
}
