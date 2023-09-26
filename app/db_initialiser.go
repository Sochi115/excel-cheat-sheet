package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Sochi115/excel-cheat-sheet/helper"
)

func initialiseDatabase(dbName string, csvFile string) {
	createDatabase(dbName)

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	createTable(db)

	csvData := helper.ReadCsvFile(csvFile)

	for _, csvLine := range csvData {
		addExcelCommandEntry(db, csvLine)
	}
}

func createDatabase(dbName string) {
	file, err := os.Create(dbName)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

func createTable(db *sql.DB) {
	excel_table_query := `CREATE TABLE excel_commands(
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "Function" TEXT,
    "Description" TEXT,
    "Syntax" TEXT,
    "Tag" TEXT,
    "Long" TEXT
  );`

	query, err := db.Prepare(excel_table_query)
	if err != nil {
		log.Fatal(err)
	}

	query.Exec()

	fmt.Println("TABLE CREATED")
}

func addExcelCommandEntry(db *sql.DB, data []string) {
	query := `INSERT INTO excel_commands(Function, Description, Syntax, Tag, Long) VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	function := data[0]
	desc := data[1]
	syntax := data[2]
	tag := data[3]
	long := data[4]

	_, err = stmt.Exec(function, desc, syntax, tag, long)

	if err != nil {
		log.Fatal(err)
	}
}
