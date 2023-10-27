package routes

import (
	"jora/app/http/controllers/attendanceController"
	"jora/app/http/controllers/auth"
	"jora/app/http/middleware"
	"jora/utility"
	"net/http"

	
	"github.com/gin-gonic/gin"
)

func Register() {
	r := app

	r.Use(gin.Logger())

	// add api prefix
	api := r.Group("/api")

	// r.Use(middleware.TrimMiddleware()) // todo

	// ping
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

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

	if (utility.Getenv("SERVE_MODE", "naturel") == "naturel") {
		r.Run(":8181")
	}

}

// @ vercel
var (
	app *gin.Engine
)

// @ vercel
func init() {
	app = gin.New()
}

// @ vercel
func Handler(w http.ResponseWriter, r *http.Request) {

	gin.SetMode(gin.ReleaseMode)
	Register()
	app.ServeHTTP(w, r)
}
