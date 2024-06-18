package cmd

import (
	"metro/data"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise a new metro database and table",
	Long:  `Initialise a new metro database and table`,
	Run: func(cmd *cobra.Command, args []string) {
		data.CreateTable()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
