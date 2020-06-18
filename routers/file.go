package routers

import (
	"github.com/gin-gonic/gin"
	c "github.com/mhdiiilham/ginorm/controllers"
)

// File ...
func File(r *gin.Engine) {

	r.POST("/image", c.UploadSingleImage) // Upload single image

}
