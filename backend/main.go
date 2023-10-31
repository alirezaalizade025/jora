package main

import (
	"jora/api"
	attendanceModel "jora/app/models/attendance"
	teamLeadModel "jora/app/models/teamLead"
	"jora/app/models"
	"jora/database/postgres"
	"jora/utility"
	"time"
	_ "time/tzdata"
	// "jora/cmd"
)

func init() {

	// set pubic/private tokens in memory
	utility.GetTokens()
}

func main() {

		// load env
		utility.LoadDotEnv()

		// run commands
		// cmd.Execute()


	// set global time zone
	// todo: move this to a better place
	loc, _ := time.LoadLocation("Asia/Tehran")
    // handle err
    time.Local = loc // -> this is setting the global timezone

	// establish db connection & migration
	postgres.ConnectDataBase()

	// migrations
	postgres.DB.AutoMigrate(&model.User{})
	postgres.DB.AutoMigrate(&utility.TokenDetails{})
	postgres.DB.AutoMigrate(&attendanceModel.Attendance{})
	postgres.DB.AutoMigrate(&teamLeadModel.TeamLead{})
	postgres.DB.AutoMigrate(&model.Company{})
	postgres.DB.AutoMigrate(&model.Team{})

	routes.Register()
}
