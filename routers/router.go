package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mhdiiilham/ginorm/middlewares"
)

// Router main func
func Router() *gin.Engine {
	r := gin.Default()

	User(r)
	r.Use(middleware.Authentication()) 
	Todo(r)

	return r
}