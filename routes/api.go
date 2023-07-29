package routes

import (
	"github.com/gin-gonic/gin"

	"jora/app/http/controllers/attendanceController"
	"jora/app/http/controllers/auth"
	"jora/app/http/middleware"
)

func Register() {
	r := gin.New()

	// r.Use(middleware.TrimMiddleware()) // todo

	// authentication
	r.POST("/login", auth.Login) // todo: if user logged in redirect
	r.POST("/logout", auth.Logout)

	// clockwise
	clockwiseGroup := r.Group("/attendances").Use(middleware.JwtAuthMiddleware())

	clockwiseGroup.POST("/clock-in", attendanceController.ClockIn)
	clockwiseGroup.POST("/clock-out", attendanceController.ClockOut)

	clockwiseGroup.POST("/leave", attendanceController.Leave)

	r.Run(":8181")
}
