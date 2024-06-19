package cmd

import (
	"metro/data"
	"metro/handlers"

	"github.com/spf13/cobra"
)

func createNewStation() {
	namePromptContent := handlers.PromptContent{
		ErrorMsg: "Введите имя станции: ",
		Label:    "Введите имя станции: ",
	}

	name := handlers.PromptGetInput(namePromptContent)

	data.InsertStations(name)
}

var newStationCmd = &cobra.Command{
	Use:   "new",
	Short: "Добавление новой станции",
	Long:  `Добавление новой станции`,
	Run: func(cmd *cobra.Command, args []string) {
		createNewStation()
	},
}

func init() {
	stationCmd.AddCommand(newStationCmd)
}
