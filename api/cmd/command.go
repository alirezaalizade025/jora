package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Jora",
	Short: "Jora CLI program",
}


func Execute() {
	rootCmd.AddCommand(createToken)
	rootCmd.Execute()
}