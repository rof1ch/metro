package cmd

import (
	"fmt"
	"metro/data"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var listRoleCmd = &cobra.Command{
	Use:   "list",
	Short: "Получить список ролей",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		roles := data.GetRoles()

		t := table.NewWriter()

		t.AppendHeader(table.Row{"ID", "Name"})

		for _, role := range roles {
			t.AppendRow(table.Row{role.Id, role.Name})
		}
		fmt.Println(t.Render())
	},
}

func init() {
	roleCmd.AddCommand(listRoleCmd)
}
