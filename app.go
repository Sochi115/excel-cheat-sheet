package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
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
	http.HandleFunc("/", a.mainPage)
	http.HandleFunc("/search", a.searchQuery)
	log.Fatal(http.ListenAndServe(":1211", nil))
}

func (a *App) mainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	tmpl.Execute(w, nil)
}

func (a *App) searchQuery(w http.ResponseWriter, r *http.Request) {
	// result := r.URL.Query().Get("q")
	result := a.getTenEntriesFromDb()
	tmpl, err := template.ParseFiles("templates/results.html")
	if err != nil {
		log.Fatal(err)
	}
	// tmpl.Execute(os.Stdout, result)
	tmpl.Execute(w, result)
}
