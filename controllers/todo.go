package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mhdiiilham/ginorm/db"
	m "github.com/mhdiiilham/ginorm/models"
	log "github.com/sirupsen/logrus"
)

// GetMyTodo ...
func GetMyTodo(c *gin.Context) {
	// Decalre variable that needed
	var todos []m.Todo

	// Fetch all todo from DB
	// and assign it to todos
	if err := db.MySQL().Where("user_id = ?", c.MustGet("userID")).Find(&todos); err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error. Fething Todo failed, please try again!"})
	}

	// If all okay, send todo to client
	c.JSON(200, gin.H{
		"message": "Fetching todos success",
		"data":    todos,
	})
}

// CreateTodo ...
func CreateTodo(c *gin.Context) {
	// Delcare all variables that needed
	userID := fmt.Sprintf("%v", c.MustGet("userID"))
	var body m.TodoInput

	// Validate the user input
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"errors": err.Error()})
		return
	}

	// Create a todo
	todo := m.Todo{
		Title:  body.Title,
		UserID: userID,
	}

	// Save Todo
	if err := db.MySQL().Save(&todo); err != nil {
		c.JSON(500, gin.H{"errors": "Saving todo failed. Please try again"})
		return
	}

	// Return todo if success
	c.JSON(200, gin.H{"todo": todo})
}

// GetTodoByID ...
func GetTodoByID(c *gin.Context) {
	// Declare variable needed
	var todo m.Todo

	// Find Todo that user look for
	// and assign it to `todo` variable
	if err := db.MySQL().Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		log.Warn("THERE ERROR ON FECTHING TODO BY IT ID");
		c.JSON(500, gin.H{"errors": "Internal Server Error"})
		return
	}
	// Check if that todo
	// belong to the user that signed in
	// If not, return 400 http status
	if isValid := todo.UserID == fmt.Sprintf("%v", c.MustGet("userID")); !isValid {
		c.JSON(400, gin.H{"message": "You Cannot Access Someone Else's Todo!", "data": nil})
		return
	}

	// IF everything is okay, then send the todo
	c.JSON(200, gin.H{
		"Message": "Fetching Todo success",
		"data": todo,
	})
}
