package models

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	FirstName string `json:"firstname" gorm:"type:varchar(50)"`
	LastName string `json:"lastname" gorm:"type:varchar(50)"`
	Email string `json:"email" gorm:"type:varchar(100);unique_index"`
	Todos []Todo `json:"todos" gorm:"foreignKey:UserID"`
	PasswordHashed string `json:"password_hashed" gorm:"type:varchar(60)"`
}

// UserSignUpInput ...
type UserSignUpInput struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName string `json:"lastname" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserLoginInput ...
type UserLoginInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
