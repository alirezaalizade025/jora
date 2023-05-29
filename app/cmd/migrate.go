package cmd

import (
	"nomasho/database/postgres"
	"nomasho/utility"
	
	"github.com/spf13/cobra"
	"gorm.io/gorm"

	userModal "nomasho/app/models/user"
)

// command to do auto migrate by gorm
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Auto migrate database",
	Long:  `Auto migrate database`,
	Run: func(cmd *cobra.Command, args []string) {
		Migration()
	},
}

func init()  {
	rootCmd.AddCommand(migrateCmd)
}


func Migration() {

	appMode := utility.Getenv("APP_MODE", "development")

	var db *gorm.DB
	if appMode == "test" {
		dbname := utility.Getenv("DB_NAME", "nomasho_db_test")
		postgres.CreateDatabase(dbname)
		db = postgres.TestConnection()
	} else {
		dbname := utility.Getenv("DB_NAME", "nomasho_db")
		postgres.CreateDatabase(dbname)
		db = postgres.Connection().Conn
	}

	// auto migrate all models
	db.AutoMigrate(&userModal.User{})
}
