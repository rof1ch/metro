package cmd

import (
	"github.com/spf13/cobra"
)

// stationVisitCmd represents the stationVisit command
var stationVisitCmd = &cobra.Command{
	Use: "stationVisit",
	Short: `- [new] Добавление записей о посещение станций
	       - [list] Вывод списка записей о посещение станций`,
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(stationVisitCmd)
}
