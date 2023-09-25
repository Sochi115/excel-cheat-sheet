package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type App struct {
	DB *sql.DB
}

func (a *App) initializeApp(dbFile string, csvFile string) {
	_, err := os.Stat(dbFile)
	if err != nil {
		initialiseDatabase(dbFile, csvFile)
	}
	a.DB, err = sql.Open("sqlite3", dbFile)

	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) run() {
	fs := http.FileServer(http.Dir("templates/excel_images"))
	port := ":1211"

	http.Handle("/excel_images/", http.StripPrefix("/excel_images", fs))
	http.HandleFunc("/", a.mainPage)
	http.HandleFunc("/search", a.searchQuery)

	fmt.Println("Serving on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func (a *App) mainPage(w http.ResponseWriter, r *http.Request) {
	pages := r.URL.Path[1:]
	// Default page
	if len(pages) == 0 {
		initial_list := a.getTenEntriesFromDb()
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
		result := a.getTenEntriesFromDb()
		tmpl.Execute(w, result)
	} else {

		commands1 := a.getFunctionsContaining(query)
		commands2 := a.getDescriptionsContaining(query)
		commands3 := a.getLongDescriptionsContaining(query)

		result := a.combineQueryResults(commands1, commands2, commands3)
		tmpl.Execute(w, result)
	}
}

func (a *App) combineQueryResults(queryResults ...[]ExcelCommand) []ExcelCommand {
	var combinedSlice []ExcelCommand
	allKeys := make(map[string]bool)
	results := []ExcelCommand{}

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

func (a *App) handleDetailsPage(w http.ResponseWriter, function_name string) {
	result := a.getByFunction(function_name)

	tmpl, _ := template.ParseFiles("templates/details.html")
	tmpl.Execute(w, result)
}
