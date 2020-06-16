package models

import "github.com/jinzhu/gorm"

// Todo model
type Todo struct {
	gorm.Model
	Title string `json:"title" gorm:"type:varchar(100)"`
	UserID uint
}