package routers

import (
	"github.com/gin-gonic/gin"
	c "github.com/mhdiiilham/ginorm/controllers"
)

// Todo ...
func Todo(r *gin.Engine) {

	r.GET("/todos", c.GetMyTodo) // Get all User's todo
	r.GET("/todos/:id", c.GetTodoByID) // Get one of user's todo
	r.POST("/todos", c.CreateTodo) // Creata a todo for user

}
