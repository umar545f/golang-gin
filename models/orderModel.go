package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	TableID     uint        `json:"table_id" gorm:"not null"`
	TotalAmount float64     `json:"total_amount" gorm:"not null"`
	Status      string      `json:"status" gorm:"size:50"` // e.g., "Pending", "Completed"
	OrderDate   time.Time   `json:"order_date" gorm:"default:CURRENT_TIMESTAMP"`
	OrderItems  []OrderItem `json:"order_items"`
}
