package cmd

import (
	"metro/data"
	"metro/handlers"

	"github.com/spf13/cobra"
)

func createNewRole() {
	namePromptContent := handlers.PromptContent{
		ErrorMsg: "Введите имя роли",
		Label:    "Как вы хотите назвать роль? ",
	}

	name := handlers.PromptGetInput(namePromptContent)

	data.InsertRole(name)
}

var newRoleCmd = &cobra.Command{
	Use:   "new",
	Short: "Добавление новой роли",
	Long:  `Добавление новой роли`,
	Run: func(cmd *cobra.Command, args []string) {
		createNewRole()
	},
}

func init() {
	roleCmd.AddCommand(newRoleCmd)
}
