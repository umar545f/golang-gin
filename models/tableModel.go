package models

import "github.com/jinzhu/gorm"

type Table struct {
	gorm.Model
	TableNumber int    `json:"table_number" gorm:"size:50;not null"`
	Seats       int    `json:"seats" gorm:"not null"`
	Location    string `json:"location" gorm:"size:100"`
	IsOccupied  bool   `json:"is_occupied"`
}
