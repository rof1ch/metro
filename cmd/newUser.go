package cmd

import (
	"metro/data"
	"metro/handlers"

	"github.com/spf13/cobra"
)

func createNewUser() {
	namePromptContent := handlers.PromptContent{
		ErrorMsg: "Вы ввели неправильно имя, введите корректно ",
		Label:    "Введите имя для пользователя: ",
	}
	name := handlers.PromptGetInput(namePromptContent)

	loginPromptContent := handlers.PromptContent{
		ErrorMsg: "Вы ввели неправильно логин, введите корректно ",
		Label:    "Введите логин пользователя: ",
	}

	login := handlers.PromptGetInput(loginPromptContent)

	passwordPromptContent := handlers.PromptContent{
		ErrorMsg: "Неправильно введен пароль",
		Label:    "Введите пароль пользователя: ",
	}

	password := handlers.PromptGetInput(passwordPromptContent)

	roles := data.GetNamesRoles()

	rolePromptContent := handlers.PromptContent{
		ErrorMsg: "Не корректная роль",
		Label:    "Выберите роль пользователя: ",
	}

	role := handlers.PromptGetSelect(rolePromptContent, roles)

	data.InsertUser(name, login, password, role)
}

var newUserCmd = &cobra.Command{
	Use:   "new",
	Short: "Добавление пользователя",
	Long:  `Добавление пользователя`,
	Run: func(cmd *cobra.Command, args []string) {
		createNewUser()
	},
}

func init() {
	userCmd.AddCommand(newUserCmd)
}
