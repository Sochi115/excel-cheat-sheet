package app

import (
	"database/sql"
	"log"
	"strings"

	"github.com/Sochi115/excel-cheat-sheet/models"
)

func (a *App) getTenEntriesFromDb() []models.ExcelCommand {
	queried_commands := []models.ExcelCommand{}
	query := "SELECT * FROM excel_commands LIMIT 10;"

	rows, err := a.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var ec models.ExcelCommand
		err = rows.Scan(&ec.Id, &ec.Function, &ec.Desc, &ec.Syntax, &ec.Tag, &ec.Long)

		if err != nil {
			queried_commands = append(queried_commands, models.ExcelCommand{})
		}

		queried_commands = append(queried_commands, ec)
	}

	return queried_commands
}

func (a *App) getByFunction(function_string string) models.ExcelCommand {
	q := `SELECT * FROM excel_commands AS e WHERE LOWER(e.Function) = ?`
	var ec models.ExcelCommand

	command_row := a.DB.QueryRow(q, strings.ToLower(function_string))

	switch err := command_row.Scan(&ec.Id, &ec.Function, &ec.Desc, &ec.Syntax, &ec.Tag, &ec.Long); err {
	case sql.ErrNoRows:
		return models.ExcelCommand{}
	case nil:
		return ec
	default:
		panic(err)
	}
}

func (a *App) getFunctionsContaining(function_string string) []models.ExcelCommand {
	q := `SELECT * FROM excel_commands AS e WHERE LOWER(e.Function) LIKE '%' || $1 || '%'`

	command_rows, err := a.DB.Query(q, strings.ToLower(function_string))
	if err != nil {
		log.Fatal(err)
	}

	if command_rows == nil {
		return nil
	}

	queried_commands := []models.ExcelCommand{}

	for command_rows.Next() {
		var ec models.ExcelCommand
		err = command_rows.Scan(&ec.Id, &ec.Function, &ec.Desc, &ec.Syntax, &ec.Tag, &ec.Long)

		if err != nil {
			queried_commands = append(queried_commands, models.ExcelCommand{})
		}

		queried_commands = append(queried_commands, ec)
	}

	return queried_commands
}

func (a *App) getDescriptionsContaining(function_string string) []models.ExcelCommand {
	q := `SELECT * FROM excel_commands AS e WHERE LOWER(e.Description) LIKE '%' || $1 || '%'`

	command_rows, err := a.DB.Query(q, strings.ToLower(function_string))
	if err != nil {
		log.Fatal(err)
	}

	if command_rows == nil {
		return nil
	}

	queried_commands := []models.ExcelCommand{}

	for command_rows.Next() {
		var ec models.ExcelCommand
		err = command_rows.Scan(&ec.Id, &ec.Function, &ec.Desc, &ec.Syntax, &ec.Tag, &ec.Long)

		if err != nil {
			queried_commands = append(queried_commands, models.ExcelCommand{})
		}

		queried_commands = append(queried_commands, ec)
	}

	return queried_commands
}

func (a *App) getLongDescriptionsContaining(function_string string) []models.ExcelCommand {
	q := `SELECT * FROM excel_commands AS e WHERE LOWER(e.Long) LIKE '%' || $1 || '%'`

	command_rows, err := a.DB.Query(q, strings.ToLower(function_string))
	if err != nil {
		log.Fatal(err)
	}

	if command_rows == nil {
		return nil
	}

	queried_commands := []models.ExcelCommand{}

	for command_rows.Next() {
		var ec models.ExcelCommand
		err = command_rows.Scan(&ec.Id, &ec.Function, &ec.Desc, &ec.Syntax, &ec.Tag, &ec.Long)

		if err != nil {
			queried_commands = append(queried_commands, models.ExcelCommand{})
		}

		queried_commands = append(queried_commands, ec)
	}

	return queried_commands
}

func (a *App) combineQueryResults(queryResults ...[]models.ExcelCommand) []models.ExcelCommand {
	var combinedSlice []models.ExcelCommand
	allKeys := make(map[string]bool)
	results := []models.ExcelCommand{}

	for _, q := range queryResults {
		combinedSlice = append(combinedSlice, q...)
	}

	for _, item := range combinedSlice {
		if _, value := allKeys[item.Function]; !value {
			allKeys[item.Function] = true
			results = append(results, item)
		}
	}

	return results
}
