package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
)

type App struct {
	DB *sql.DB
}

func (a *App) InitializeApp(dbFile string, csvFile string) {
	_, err := os.Stat(dbFile)
	if err != nil {
		initialiseDatabase(dbFile, csvFile)
	}
	a.DB, err = sql.Open("sqlite3", dbFile)

	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) Run() {
	fs := http.FileServer(http.Dir("templates/excel_images"))
	port := ":1211"

	http.Handle("/excel_images/", http.StripPrefix("/excel_images", fs))
	http.HandleFunc("/", a.mainPage)
	http.HandleFunc("/search", a.searchQuery)

	fmt.Println("Serving on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
