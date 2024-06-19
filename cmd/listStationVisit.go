package cmd

import (
	"fmt"
	"metro/data"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// listStationVisitCmd represents the listStationVisit command
var listStationVisitCmd = &cobra.Command{
	Use:   "list",
	Short: "Вывести список посещений станций",
	Long:  `Вывести список посещений станций`,
	Run: func(cmd *cobra.Command, args []string) {
		stationsVisits := data.GetStationVisits()

		t := table.NewWriter()

		t.AppendHeader(table.Row{"ID", "Station", "Datetime"})

		for _, val := range stationsVisits {
			t.AppendRow(table.Row{val.Id, val.Station_name, val.Datetime})
		}

		fmt.Println(t.Render())
	},
}

func init() {
	stationVisitCmd.AddCommand(listStationVisitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listStationVisitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listStationVisitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
