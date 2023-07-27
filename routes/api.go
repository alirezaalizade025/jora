package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"jora/app/middleware"
	"jora/controllers/auth"
)

func Register() {
	r := gin.Default()

	// ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}, middleware.JwtAuthMiddleware())

	// authentication
	r.POST("/login", auth.Login)
	r.POST("/logout", auth.Logout)

	r.Run(":8181")
}
