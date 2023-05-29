package main

import (
	"nomasho/routes"
	"nomasho/utility"
	"nomasho/app/cmd"
)

func main() {

	utility.LoadDotEnv()

	cmd.Execute()

    routes.Register()
}