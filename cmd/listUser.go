package cmd

import (
	"fmt"
	"metro/data"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var listUserCmd = &cobra.Command{
	Use:   "list",
	Short: "Список пользователей",
	Long: `Список пользователей`,
	Run: func(cmd *cobra.Command, args []string) {
		users := data.GetUsers()

		t := table.NewWriter()

		t.AppendHeader(table.Row{"ID", "Name", "Login", "Password", "Role"})

		for _, user := range users{
			t.AppendRow(table.Row{user.Id, user.Name, user.Login, user.Password, user.Role_name})
		}

		fmt.Println(t.Render())
	},
}

func init() {
	userCmd.AddCommand(listUserCmd)
}
