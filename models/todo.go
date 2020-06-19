package models

import "github.com/jinzhu/gorm"

// Todo model
type Todo struct {
	gorm.Model
	Title  string `json:"title" gorm:"type:varchar(100)"`
	UserID string `json:"user_id" gorm:"type:varchar(10)"`
}

// TodoInput ...
type TodoInput struct {
	Title string
}
