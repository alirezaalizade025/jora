package main

import (
	"nomasho/database/postgres"
	"nomasho/routes"
	"nomasho/utility"
)

func main() {

	utility.LoadDotEnv()

	// establish db connection & migration
	postgres.ConnectDataBase()

    routes.Register()
}