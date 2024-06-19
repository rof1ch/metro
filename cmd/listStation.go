package cmd

import (
	"fmt"
	"metro/data"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var listStationCmd = &cobra.Command{
	Use:   "list",
	Short: "Вывести список станций",
	Long:  `Вывести список станций`,
	Run: func(cmd *cobra.Command, args []string) {
		stations := data.GetStations()

		t := table.NewWriter()

		t.AppendHeader(table.Row{"ID", "Name"})

		for _, station := range stations {
			t.AppendRow(table.Row{station.Id, station.Name})
		}

		fmt.Println(t.Render())
	},
}

func init() {
	stationCmd.AddCommand(listStationCmd)
}
