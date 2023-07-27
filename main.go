package main

import (
	userModel "jora/app/models/user"
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

	routes.Register()
}
