package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	// Todo: Add a show template here
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func main() {
	port := ":3333"
	http.HandleFunc("/", mainPage)
	fmt.Println("Server Listening on port ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
