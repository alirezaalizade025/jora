package main

import (
	teamLeadModel "jora/app/models/teamLead"
	attendanceModel "jora/app/models/attendance"
	userModel "jora/app/models/user"
	"jora/database/postgres"
	"jora/routes"
	"jora/utility"
	"time"
	_ "time/tzdata"
	"jora/cmd"
)

func init() {

	// set pubic/private tokens in memory
	utility.GetTokens()
}

func main() {

		// load env
		utility.LoadDotEnv()

		// run commands
		cmd.Execute()


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
	postgres.DB.AutoMigrate(&teamLeadModel.TeamLead{})

	routes.Register()
}
