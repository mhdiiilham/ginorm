package routers

import (
	"github.com/gin-gonic/gin"
	c "github.com/mhdiiilham/ginorm/controllers"
)

// User auth handler
func User(r *gin.Engine) {

	r.POST("/signup", c.CreateUser) // Create new user
	r.POST("/login", c.Login) // Login existing user

}
