package main

import (
	userModel "jora/app/models/user"
	attendanceModel "jora/app/models/attendance"
	"jora/database/postgres"
	"jora/routes"
	"jora/utility"
)

func main() {

	utility.LoadDotEnv()

	// establish db connection & migration
	postgres.ConnectDataBase()

	// migrations
	postgres.DB.AutoMigrate(&userModel.User{})
	postgres.DB.AutoMigrate(&utility.TokenDetails{})
	postgres.DB.AutoMigrate(&attendanceModel.Attendance{})

	routes.Register()
}
