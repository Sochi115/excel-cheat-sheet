package main

var (
	csvFile string = "excel_commands.csv"
	dbFile  string = "ecs.db"
)

func main() {
	app := App{}

	app.initializeApp(dbFile, csvFile)
	app.run()
}

// func mainPage(w http.ResponseWriter, r *http.Request) {
// 	tmpl := template.Must(template.ParseFiles("templates/index.html"))
//
// 	tmpl.Execute(w, nil)
// }
//
// func getFilms(w http.ResponseWriter, r *http.Request) {
// 	tmpl := template.Must(template.ParseFiles("templates/db_items.html"))
//
// 	tmpl.Execute(w, nil)
// }
//
// func queryTests(w http.ResponseWriter, r *http.Request) {
// 	query := r.URL.Query().Get("q")
// 	fmt.Println(query)
// 	io.WriteString(w, query)
// }
//
// func main() {
// 	port := ":3333"
// 	http.HandleFunc("/", mainPage)
// 	http.HandleFunc("/get-movies", getFilms)
// 	http.HandleFunc("/search", queryTests)
// 	fmt.Println("Server Listening on port ", port)
// 	log.Fatal(http.ListenAndServe(port, nil))
// 	getTenEntriesFromDb()
// }
