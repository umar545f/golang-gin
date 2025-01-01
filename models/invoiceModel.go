package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Invoice struct {
	gorm.Model
	OrderID       uint      `json:"order_id" gorm:"not null"`
	Amount        float64   `json:"amount" gorm:"not null"`
	PaymentStatus string    `json:"payment_status" gorm:"size 50"` //"Paid", "Unpaid"
	InvoiceDate   time.Time `json:"invoice_date" gorm:"default:CURRENT_TIMESTAMP"`
}
