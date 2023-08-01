package routes

import (
	"jora/app/http/controllers/attendanceController"
	"jora/app/http/controllers/auth"
	"jora/app/http/middleware"

	"github.com/gin-gonic/gin"
)

func Register() {
	r := gin.New()

	// r.Use(middleware.TrimMiddleware()) // todo

	// authentication
	r.POST("/login", auth.Login) // todo: if user logged in redirect
	r.POST("/logout", auth.Logout)

	// clockwise
	clockwiseGroup := r.Group("/attendances").Use(middleware.JwtAuthMiddleware())

	clockwiseGroup.POST("/clock-in", attendanceController.Start)
	clockwiseGroup.POST("/clock-out", attendanceController.Finish)

	clockwiseGroup.POST("/leave", attendanceController.Leave)
	clockwiseGroup.POST("/leave/hourly", attendanceController.HourlyLeave)
	clockwiseGroup.POST("business-trip", attendanceController.BusinessTrip)
	clockwiseGroup.POST("remote-work", attendanceController.RemoteWork)
	clockwiseGroup.POST("missing", attendanceController.MissingAttendance)

	clockwiseGroup.PUT("/working/:id", attendanceController.Update)

	r.Run(":8181")
}
