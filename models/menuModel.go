package models

import "github.com/jinzhu/gorm"

// Menu represents a menu category (breakfast, lunch, dinner)  ;-)
type Menu struct {
	gorm.Model
	Name        string `json:"name" gorm:"size:100;not null"`
	Description string `json:"description" gorm:"type:text"`
	Foods       []Food `json:"foods" gorm:"many2many:menu_foods;foreignKey:ID;joinForeignKey:menu_id;References:ID;joinReferences:food_id"`
}
