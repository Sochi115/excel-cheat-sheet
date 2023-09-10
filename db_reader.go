package main

import (
	"log"
)

func (a *App) getTenEntriesFromDb() []ExcelCommand {
	queried_commands := []ExcelCommand{}
	query := "SELECT * FROM excel_commands LIMIT 10;"

	rows, err := a.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var ec ExcelCommand
		err = rows.Scan(&ec.Id, &ec.Function, &ec.Desc, &ec.Syntax, &ec.Tag)

		if err != nil {
			queried_commands = append(queried_commands, ExcelCommand{})
		}

		queried_commands = append(queried_commands, ec)
	}

	return queried_commands
}
