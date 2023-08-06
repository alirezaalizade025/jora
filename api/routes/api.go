package routes

import (
	"jora/app/http/controllers/attendanceController"
	"jora/app/http/controllers/auth"
	"jora/app/http/middleware"

	"github.com/gin-gonic/gin"
)

func Register() {
	r := gin.New()

	// add api prefix
	api := r.Group("/api")

	// r.Use(middleware.TrimMiddleware()) // todo

	// authentication
	api.POST("/login", auth.Login) // todo: if user logged in redirect
	api.POST("/logout", auth.Logout)

	// clockwise
	clockwiseGroup := api.Group("/attendances").Use(middleware.JwtAuthMiddleware())

	clockwiseGroup.POST("/clock-in", attendanceController.Start)
	clockwiseGroup.POST("/clock-out", attendanceController.Finish)

	clockwiseGroup.POST("/leave", attendanceController.Leave)
	clockwiseGroup.POST("/leave/hourly", attendanceController.HourlyLeave)
	clockwiseGroup.POST("business-trip", attendanceController.BusinessTrip)
	clockwiseGroup.POST("remote-work", attendanceController.RemoteWork)
	clockwiseGroup.POST("missing", attendanceController.MissingAttendance)

	clockwiseGroup.PUT("/working/:id", attendanceController.Update)
	// todo: team lead check

	// todo: manager check

	r.Run(":8181")
}
