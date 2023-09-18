package main

import (
	"database/sql"
	"log"
	"strings"
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
		err = rows.Scan(&ec.Id, &ec.Function, &ec.Desc, &ec.Syntax, &ec.Tag, &ec.Long)

		if err != nil {
			queried_commands = append(queried_commands, ExcelCommand{})
		}

		queried_commands = append(queried_commands, ec)
	}

	return queried_commands
}

func (a *App) getByFunction(function_string string) ExcelCommand {
	q := `SELECT * FROM excel_commands AS e WHERE LOWER(e.Function) = ?`
	var ec ExcelCommand

	command_row := a.DB.QueryRow(q, strings.ToLower(function_string))

	switch err := command_row.Scan(&ec.Id, &ec.Function, &ec.Desc, &ec.Syntax, &ec.Tag, &ec.Long); err {
	case sql.ErrNoRows:
		return ExcelCommand{}
	case nil:
		return ec
	default:
		panic(err)
	}
}

func (a *App) getFunctionsContaining(function_string string) []ExcelCommand {
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
		err = command_rows.Scan(&ec.Id, &ec.Function, &ec.Desc, &ec.Syntax, &ec.Tag, &ec.Long)

		if err != nil {
			queried_commands = append(queried_commands, ExcelCommand{})
		}

		queried_commands = append(queried_commands, ec)
	}

	return queried_commands
}

func (a *App) getDescriptionsContaining(function_string string) []ExcelCommand {
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
		err = command_rows.Scan(&ec.Id, &ec.Function, &ec.Desc, &ec.Syntax, &ec.Tag, &ec.Long)

		if err != nil {
			queried_commands = append(queried_commands, ExcelCommand{})
		}

		queried_commands = append(queried_commands, ec)
	}

	return queried_commands
}
