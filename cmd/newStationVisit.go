package cmd

import (
	"metro/data"
	"metro/handlers"

	"github.com/spf13/cobra"
)

func createNewStationVisit() {
	stationPromptContent := handlers.PromptContent{
		ErrorMsg: "Выберите имя станции: ",
		Label:    "Выберите имя станции: ",
	}

	stations := data.GetStationsName()

	station_name := handlers.PromptGetSelect(stationPromptContent, stations)

	station_id := data.GetStationIdByName(station_name)

	data.InsertStationVisit(station_id)
}

var newStationVisitCmd = &cobra.Command{
	Use:   "new",
	Short: "Добавление записи посещения станции",
	Long:  `Добавление записи посещения станции`,
	Run: func(cmd *cobra.Command, args []string) {
		createNewStationVisit()
	},
}

func init() {
	stationVisitCmd.AddCommand(newStationVisitCmd)
}
