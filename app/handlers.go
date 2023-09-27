package app

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
	"time"
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
	start := time.Now()
	query := r.URL.Query().Get("q")

	tmpl, err := template.ParseFiles("templates/results.html")
	if err != nil {
		log.Fatal(err)
	}

	if len(query) == 0 {
		result := a.getDefaultTenCommands()
		tmpl.Execute(w, result)
	} else {

		commands1 := a.getWeightedFunctionsContaining(query)
		commands2 := a.getWeightedDescriptionsContaining(query)
		commands3 := a.getWeightedLongDescriptionsContaining(query)

		result := a.combineQueryResults(commands1, commands2, commands3)
		tmpl.Execute(w, result)
		fmt.Println("Time taken:", time.Since(start))
	}
}

func (a *App) handleDetailsPage(w http.ResponseWriter, function_name string) {
	result := a.getByFunction(function_name)

	tmpl, _ := template.ParseFiles("templates/details.html")
	tmpl.Execute(w, result)
}
