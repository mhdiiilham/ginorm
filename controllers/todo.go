package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mhdiiilham/ginorm/db"
	m "github.com/mhdiiilham/ginorm/models"
	log "github.com/sirupsen/logrus"
)

// GetMyTodo ...
func GetMyTodo(c *gin.Context) {
	var todos []m.Todo
	log.Info("tetap masuk sini ga sih?")
	db.MySQL().Where("user_id = ?", c.MustGet("userID")).Find(&todos)
	c.JSON(200, gin.H{
		"message": "Fetching todos success",
		"data":    todos,
	})
}

// CreateTodo ...
func CreateTodo(c *gin.Context) {
	var body m.TodoInput

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"errors": err.Error()})
		return
	}
	user := m.User{}
	getUser := db.MySQL().Where("id = ?", c.MustGet("userID")).Find(&user)
	if getUser.Error != nil {
		c.JSON(500, gin.H{"errors": "Internal Server Error"})
		return
	}
	todo := m.Todo{
		Title:  body.Title,
		UserID: user.ID,
	}
	saving := db.MySQL().Save(&todo)
	if saving.Error != nil {
		c.JSON(500, gin.H{"errors": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"todo": todo})
}

// GetTodoByID ...
func GetTodoByID(c *gin.Context) {
	var todo m.Todo
	if err := db.MySQL().Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		log.Warn("THERE ERROR ON FECTHING TODO BY IT ID");
		c.JSON(500, gin.H{"errors": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{
		"Message": "Fetching Todo success",
		"data": todo,
	})
}
