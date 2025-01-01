package models

import "github.com/jinzhu/gorm"

// Food represents a food item in the menu :-)
type Food struct {
	gorm.Model
	Name        string `json:"name" gorm:"size:255;not null"`
	Description string `json:"description" gorm:"type:text"`
	Price       int32  `json:"price" gorm:"not null"`
	Category    string `json:"category" gorm:"size:100"`
	OnMenu      bool   `json:"default:false"`
}
