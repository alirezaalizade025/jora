package main

import (
	attendanceModel "jora/app/models/attendance"
	userModel "jora/app/models/user"
	"jora/database/postgres"
	"jora/routes"
	"jora/utility"
	"time"
	_ "time/tzdata"
)

func main() {

	utility.LoadDotEnv()

	// set global time zone
	// todo: move this to a better place
	loc, _ := time.LoadLocation("Asia/Tehran")
    // handle err
    time.Local = loc // -> this is setting the global timezone

	// establish db connection & migration
	postgres.ConnectDataBase()

	// migrations
	postgres.DB.AutoMigrate(&userModel.User{})
	postgres.DB.AutoMigrate(&utility.TokenDetails{})
	postgres.DB.AutoMigrate(&attendanceModel.Attendance{})

	routes.Register()
}
