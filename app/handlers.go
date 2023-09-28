package app

import (
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/Sochi115/excel-cheat-sheet/models"
)

func (a *App) mainPage(w http.ResponseWriter, r *http.Request) {
	pages := r.URL.Path[1:]
	// Default page
	if len(pages) == 0 {
		initial_list := a.getDefaultTenCommands()
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, initial_list)

		// Dynamic Page
	} else {
		// Get function name from request
		image_name := strings.Replace(r.URL.RequestURI(), "/", "", -1)
		a.handleDetailsPage(w, image_name)

	}
}

func (a *App) searchQuery(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	tmpl, err := template.ParseFiles("templates/results.html")
	if err != nil {
		log.Fatal(err)
	}

	if len(query) == 0 {
		result := a.getDefaultTenCommands()
		tmpl.Execute(w, result)
	} else {

		tokens := strings.Fields(query)

		queryResults := [][]models.WeightedQuery{}

		for _, t := range tokens {
			queryResults = append(queryResults, a.getWeightedFunctionsContaining(t))
			queryResults = append(queryResults, a.getWeightedFunctionEquals(t))
			queryResults = append(queryResults, a.getWeightedDescriptionsContaining(t))
			queryResults = append(queryResults, a.getWeightedLongDescriptionsContaining(t))
		}

		result := a.combineQueryResults(queryResults...)

		if len(result) == 0 {
			ec := models.ExcelCommand{}
			ec.Function = "ERROR"
			ec.Desc = "Function " + query + " was not found"
			ec.Syntax = ""
			result = append(result, ec)
		}
		tmpl.Execute(w, result)
	}
}

func (a *App) handleDetailsPage(w http.ResponseWriter, function_name string) {
	result := a.getByFunction(function_name)

	tmpl, _ := template.ParseFiles("templates/details.html")
	tmpl.Execute(w, result)
}
