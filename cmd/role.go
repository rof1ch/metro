package cmd

import (
	"github.com/spf13/cobra"
)

var roleCmd = &cobra.Command{
	Use: "role",
	Short: `- [new] Добавление роли 
	       - [list] Вывод списка ролей`,
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(roleCmd)
}
