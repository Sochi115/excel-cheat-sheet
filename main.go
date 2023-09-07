package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	tmpl.Execute(w, nil)
}

func getFilms(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/db_items.html"))

	films := map[string][]Film{
		"Films": {
			{Title: "The Godfather", Director: "Francis Ford Coppola"},
			{Title: "Blade Runner", Director: "Ridley Scott"},
			{Title: "The Thing", Director: "John Carpenter"},
		},
	}
	tmpl.Execute(w, films)
}

func main() {
	port := ":3333"
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/get-movies", getFilms)
	fmt.Println("Server Listening on port ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
