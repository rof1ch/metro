package main

import (
	"metro/cmd"
	"metro/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
