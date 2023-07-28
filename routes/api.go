package routes

import (
	"github.com/gin-gonic/gin"

	"jora/app/http/middleware"
	"jora/app/http/controllers/auth"
	"jora/app/http/controllers/clockwork"
)

func Register() {
	r := gin.Default()


	// authentication
	r.POST("/login", auth.Login)
	r.POST("/logout", auth.Logout)

	// clockwise
	clockwiseGroup := r.Group("/clockworks").Use(middleware.JwtAuthMiddleware())
	
	clockwiseGroup.POST("/clock-in", clockwork.ClockIn)
	clockwiseGroup.POST("/clock-out", clockwork.ClockOut)

	r.Run(":8181")
}
