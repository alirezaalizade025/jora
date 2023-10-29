package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func TrimMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Trim the values of form parameters
		c.Request.ParseForm()
		for key, values := range c.Request.PostForm {
			for i, value := range values {
				trimmedValue := strings.TrimSpace(value)
				values[i] = trimmedValue
			}
			c.Request.PostForm.Set(key, strings.Join(values, ","))
		}

		// Trim the values of JSON request bodies
		var jsonBody map[string]interface{}
		err := c.ShouldBindJSON(&jsonBody)
		if err == nil {
			for key, value := range jsonBody {
				if str, ok := value.(string); ok {
					jsonBody[key] = strings.TrimSpace(str)
				}
			}
			c.Set("jsonBody", jsonBody)
		}

		// Call the next middleware or handler
		c.Next()

	}
}
