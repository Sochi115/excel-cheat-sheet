package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"
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
	http.Handle("/excel_images/", http.StripPrefix("/excel_images", fs))
	http.HandleFunc("/", a.mainPage)
	http.HandleFunc("/search", a.searchQuery)
	log.Fatal(http.ListenAndServe(":1211", nil))
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

		tmpl, _ := template.ParseFiles("templates/details.html")
		tmpl.Execute(w, fmt.Sprintf("/excel_images/%v.jpg", image_name))
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
		result := a.getTenEntriesFromDb()
		tmpl.Execute(w, result)
	} else {

		commands1 := a.getFunctionsContaining(query)
		commands2 := a.getDescriptionsContaining(query)

		result := a.combineQueryResults(commands1, commands2)
		tmpl.Execute(w, result)

		fmt.Println("Time taken: ", time.Since(start))
	}
}

func (a *App) combineQueryResults(
	commands1 []ExcelCommand,
	commands2 []ExcelCommand,
) []ExcelCommand {
	combined_slice := append(commands1, commands2...)
	allKeys := make(map[string]bool)
	results := []ExcelCommand{}

	for _, item := range combined_slice {
		if _, value := allKeys[item.Function]; !value {
			allKeys[item.Function] = true
			results = append(results, item)
		}
	}

	return results
}
