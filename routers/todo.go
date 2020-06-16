package routers

import (
	"github.com/gin-gonic/gin"
	c "github.com/mhdiiilham/ginorm/controllers"
)

// Todo ...
func Todo(r *gin.Engine) {
	r.GET("/todos", c.GetMyTodo)
	r.POST("/todos", c.CreateTodo)
}
