package cmd

import (
	"metro/data"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Иницилизация и создание таблиц базы данных",
	Long:  `Иницилизация и создание таблиц базы данных`,
	Run: func(cmd *cobra.Command, args []string) {
		data.CreateTable()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
