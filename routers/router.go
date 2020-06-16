package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mhdiiilham/ginorm/middlewares"
)

// Router main func
func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/", homeFunc)
	User(r)

	// Authorization here
	r.Use(middleware.Authentication())
	Todo(r)
	return r
}

func homeFunc(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Everything is gonna be okay"})
}
