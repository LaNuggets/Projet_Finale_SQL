package main

import (
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var port = ":8766"

func main() {

	css := http.FileServer(http.Dir("./static/style"))       // For add css to the html pages
	http.Handle("/style/", http.StripPrefix("/style/", css)) // For add css to the html pages

	http.HandleFunc("/accueil", Accueil)
	http.HandleFunc("/allemployees", AllEmployees)
	http.HandleFunc("/editemployee", EditEmployeePage)
	http.HandleFunc("/employees", Employees)

	fmt.Println("http://localhost" + port + "/accueil")
	http.ListenAndServe(port, nil)
}
