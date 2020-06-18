package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	m "github.com/mhdiiilham/ginorm/models"
	log "github.com/sirupsen/logrus"
)

// MySQL ...
func MySQL() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@/company?charset=utf8&parseTime=True&loc=Local")
	db.AutoMigrate(&m.User{}, &m.Todo{}, &m.Image{})
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}
	return db
}
