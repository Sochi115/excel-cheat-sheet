package main

import "github.com/Sochi115/excel-cheat-sheet/app"

var (
	csvFile string = "excel_commands.csv"
	dbFile  string = "ecs.db"
)

func main() {
	appInstance := app.App{}

	appInstance.InitializeApp(dbFile, csvFile)
	appInstance.Run()
}
