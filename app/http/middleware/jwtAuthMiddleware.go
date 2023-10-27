package middleware

import (
	"jora/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		err := utility.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token is not valid",
			})
		
			c.Abort()
			return
		}
	

		// check with db if token is valid set user to context
		err = utility.TokenCheckDb(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token is not valid",
			})
		
			c.Abort()
			return
		}

		c.Next()
	}
}

func errorHandler(c *gin.Context, err error) {
	if err == nil {
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"error": err.Error(),
	})

	c.Abort()
}
