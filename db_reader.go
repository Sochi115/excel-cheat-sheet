package main

import (
	"log"
	"strings"
	"sync"
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

func (a *App) getFunctionsContaining(function_string string, wg *sync.WaitGroup) []ExcelCommand {
	wg.Add(1)
	q := `SELECT * FROM excel_commands AS e WHERE LOWER(e.Function) LIKE '%' || $1 || '%'`

	command_rows, err := a.DB.Query(q, strings.ToLower(function_string))
	if err != nil {
		log.Fatal(err)
	}

	if command_rows == nil {
		return nil
	}

	queried_commands := []ExcelCommand{}

	for command_rows.Next() {
		var ec ExcelCommand
		err = command_rows.Scan(&ec.Id, &ec.Function, &ec.Desc, &ec.Syntax, &ec.Tag)

		if err != nil {
			queried_commands = append(queried_commands, ExcelCommand{})
		}

		queried_commands = append(queried_commands, ec)
	}
	defer wg.Done()

	return queried_commands
}

func (a *App) getDescriptionsContaining(function_string string, wg *sync.WaitGroup) []ExcelCommand {
	wg.Add(1)
	q := `SELECT * FROM excel_commands AS e WHERE LOWER(e.Description) LIKE '%' || $1 || '%'`

	command_rows, err := a.DB.Query(q, strings.ToLower(function_string))
	if err != nil {
		log.Fatal(err)
	}

	if command_rows == nil {
		return nil
	}

	queried_commands := []ExcelCommand{}

	for command_rows.Next() {
		var ec ExcelCommand
		err = command_rows.Scan(&ec.Id, &ec.Function, &ec.Desc, &ec.Syntax, &ec.Tag)

		if err != nil {
			queried_commands = append(queried_commands, ExcelCommand{})
		}

		queried_commands = append(queried_commands, ec)
	}
	defer wg.Done()

	return queried_commands
}
