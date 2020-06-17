package routers

import (
	"github.com/gin-gonic/gin"
	c "github.com/mhdiiilham/ginorm/controllers"
)

// Todo ...
func Todo(r *gin.Engine) {
	r.GET("/todos", c.GetMyTodo)
	r.GET("/todos/{id}", c.GetTodoByID)
	r.POST("/todos", c.CreateTodo)
}
