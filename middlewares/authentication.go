package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	h "github.com/mhdiiilham/ginorm/helpers"
)

// Authentication ...
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header["Authorization"]
		if len(bearerToken) <= 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": "NOT AUTHORIZED"})
			c.Abort()
			return
		}

		token := strings.Split(bearerToken[0], " ")[1]
		if len(token) <= 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": "NOT AUTHORIZED"})
			c.Abort()
			return
		}

		// Check if token valid or not
		err := h.TokenValid(token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"errors": "TOKEN NOT VALID"})
			c.Abort()
			return
		}

		// Extracted token metadata
		metaData, err := h.ExtractedJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": "TOKEN NOT VALID"})
			c.Abort()
			return
		}

		c.Set("userID", metaData.ID)
		c.Next()
	}
}
