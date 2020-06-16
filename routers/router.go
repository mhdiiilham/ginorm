package routers

import (
	"github.com/gin-gonic/gin"
)

// Router main func
func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/", homeFunc)
	User(r)

	return r
}

func homeFunc(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Everything is gonna be okay"})
}