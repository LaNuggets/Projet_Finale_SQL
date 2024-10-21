package main

import (
	Projet_Final_SQL "Projet_Final_SQL/go"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func AllEmployees(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("./static/html/allemployees.html") // Read the home page
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	// if r.Method == http.MethodPost {
	allEmployeesList := Projet_Final_SQL.GetAllEmployees(w, r)
	fmt.Println(allEmployeesList)
	// 	return
	// }

	// data := struct {
	// 	// lastName   string
	// 	// firstName  string
	// 	// phone      string
	// 	// address    string
	// 	// birthday   string
	// 	// department string
	// 	// post       string
	// 	// wage       int
	// 	// manage     string
	// 	allEmployees []Projet_Final_SQL.Employee
	// }{
	// 	allEmployees: allEmployeesList,
	// }

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
