package app

import (
	"database/sql"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/Sochi115/excel-cheat-sheet/models"
)

func (a *App) getDefaultTenCommands() []models.ExcelCommand {
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

func (a *App) getCommandByFunction(function_string string) []models.WeightedQuery {
	q := `SELECT * FROM excel_commands AS e WHERE LOWER(e.Function) = ?`
	var ec models.ExcelCommand

	queried_commands := []models.WeightedQuery{}

	command_row := a.DB.QueryRow(q, strings.ToLower(function_string))

	switch err := command_row.Scan(&ec.Id, &ec.Function, &ec.Desc, &ec.Syntax, &ec.Tag, &ec.Long); err {
	case sql.ErrNoRows:
		return nil
	case nil:
		var wq models.WeightedQuery
		wq.ExcelCommand = ec
		wq.Score = 4
		queried_commands = append(queried_commands, wq)
		return queried_commands
	default:
		panic(err)
	}
}

func (a *App) getWeightedFunctionsContaining(function_string string) []models.WeightedQuery {
	q := `SELECT * FROM excel_commands AS e WHERE LOWER(e.Function) LIKE '%' || $1 || '%'`

	command_rows, err := a.DB.Query(q, strings.ToLower(function_string))
	if err != nil {
		log.Fatal(err)
	}

	if command_rows == nil {
		return nil
	}

	queried_commands := []models.WeightedQuery{}

	for command_rows.Next() {
		var ec models.ExcelCommand
		err = command_rows.Scan(&ec.Id, &ec.Function, &ec.Desc, &ec.Syntax, &ec.Tag, &ec.Long)

		if err != nil {
			continue
		}

		var wq models.WeightedQuery
		wq.ExcelCommand = ec
		wq.Score = 2
		queried_commands = append(queried_commands, wq)
	}

	return queried_commands
}

func (a *App) getWeightedDescriptionsContaining(query_value string) []models.WeightedQuery {
	q := `SELECT * FROM excel_commands AS e WHERE LOWER(e.Description) LIKE '%' || $1 || '%'`

	command_rows, err := a.DB.Query(q, strings.ToLower(query_value))
	if err != nil {
		log.Fatal(err)
	}

	if command_rows == nil {
		return nil
	}

	queried_commands := []models.WeightedQuery{}

	for command_rows.Next() {
		var ec models.ExcelCommand
		err = command_rows.Scan(&ec.Id, &ec.Function, &ec.Desc, &ec.Syntax, &ec.Tag, &ec.Long)

		if err != nil {
			continue
		}

		var wq models.WeightedQuery
		wq.ExcelCommand = ec
		wq.Score = 1
		queried_commands = append(queried_commands, wq)
	}

	return queried_commands
}

func (a *App) getWeightedLongDescriptionsContaining(query_value string) []models.WeightedQuery {
	q := `SELECT * FROM excel_commands AS e WHERE LOWER(e.Long) LIKE '%' || $1 || '%'`

	command_rows, err := a.DB.Query(q, strings.ToLower(query_value))
	if err != nil {
		log.Fatal(err)
	}

	if command_rows == nil {
		return nil
	}

	queried_commands := []models.WeightedQuery{}

	for command_rows.Next() {
		var ec models.ExcelCommand
		err = command_rows.Scan(&ec.Id, &ec.Function, &ec.Desc, &ec.Syntax, &ec.Tag, &ec.Long)

		if err != nil {
			continue
		}

		var wq models.WeightedQuery
		wq.ExcelCommand = ec
		wq.Score = 0.5
		queried_commands = append(queried_commands, wq)
	}

	return queried_commands
}

func (a *App) combineQueryResults(queryResults ...[]models.WeightedQuery) []models.ExcelCommand {
	var combinedSlice []models.WeightedQuery
	weightScores := make(map[string]float32)
	visited := make(map[string]bool)
	results := []models.ExcelCommand{}

	for _, q := range queryResults {
		combinedSlice = append(combinedSlice, q...)
	}

	for _, wq := range combinedSlice {
		ec := wq.ExcelCommand
		score := wq.Score

		weightScores[ec.Function] += score

		if _, value := visited[ec.Function]; !value {
			visited[ec.Function] = true
			results = append(results, ec)
		}
	}

	sort.SliceStable(results, func(i, j int) bool {
		return weightScores[results[i].Function] > weightScores[results[j].Function]
	})

	fmt.Println("Weights", weightScores)

	if len(results) <= 15 {
		return results
	} else {
		return results[:15]
	}
}
