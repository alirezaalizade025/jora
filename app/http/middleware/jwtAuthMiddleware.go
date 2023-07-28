package middleware

import (
	"jora/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		errorHandler(c, utility.TokenValid(c))

		// check with db if token is valid set user to context
		errorHandler(c, utility.TokenCheckDb(c))

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
