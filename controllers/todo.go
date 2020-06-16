package controllers

import (
	// log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/mhdiiilham/ginorm/db"
	m "github.com/mhdiiilham/ginorm/models"
)

// GetMyTodo ...
func GetMyTodo(c *gin.Context) {
	var todos []m.Todo

	metaData := c.MustGet("meta-data")

	db.MySQL().Find(&todos)
	c.JSON(200, gin.H{
		"message": "Fetching todos success",
		"data": metaData,
	})
}

// CreateTodo ...