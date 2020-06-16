package middleware

import (
	"github.com/gin-gonic/gin"
	h "github.com/mhdiiilham/ginorm/helpers"
	// log "github.com/sirupsen/logrus"
	"strings"
)

// Authentication ...
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header["Authorization"]
		if len(bearerToken) <= 0 {
			c.JSON(400, gin.H{"errors": "NOT AUTHORIZED"})
			return
		}

		token := strings.Split(bearerToken[0], " ")[1]
		if len(token) <= 0 {
			c.JSON(400, gin.H{"errors": "NOT AUTHORIZED"})
			return
		}

		// Check if token valid or not
		err := h.TokenValid(token)
		if err != nil {
			c.JSON(400, gin.H{"errors": "NOT AUTHORIZED"})
			return
		}

		// Extracted token metadata
		metaData, err := h.ExtractedJWT(token)
		if err != nil {
			c.JSON(400, gin.H{"errors": "NOT AUTHORIZED"})
			return
		}

		// c.Set("meta-data", metaData)
		c.Set("userID", metaData.ID)

		c.Next()
	}
}
