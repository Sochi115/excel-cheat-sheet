package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/savioxavier/termlink"
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

	link := "http://localhost" + port

	http.Handle("/excel_images/", http.StripPrefix("/excel_images", fs))
	http.HandleFunc("/", a.mainPage)
	http.HandleFunc("/search", a.searchQuery)

	fmt.Println(termlink.Link("Click here", link))
	log.Fatal(http.ListenAndServe(port, nil))
}
