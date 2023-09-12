package main

var (
	csvFile string = "excel_commands.csv"
	dbFile  string = "ecs.db"
)

func main() {
	app := App{}

	app.initializeApp(dbFile, csvFile)
	app.run()
}
