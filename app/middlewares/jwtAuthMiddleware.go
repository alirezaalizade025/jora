package middlewares

import (
	"net/http"
	"nomasho/utility"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utility.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized" + err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}