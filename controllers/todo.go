package controllers

import (
	// log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/mhdiiilham/ginorm/db"
	// h "github.com/mhdiiilham/ginorm/helpers"
	m "github.com/mhdiiilham/ginorm/models"
)

// GetMyTodo ...
func GetMyTodo(c *gin.Context) {
	var todos []m.Todo

	db.MySQL().Find(&todos)

	c.JSON(200, gin.H{
		"message": "Fetching todos success",
		"data": todos,
	})
}