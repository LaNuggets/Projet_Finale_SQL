package main

import (
	Projet_Final_SQL "Projet_Final_SQL/go"
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

	allEmployeesList := Projet_Final_SQL.GetAllEmployees(w, r)

	department, err := Projet_Final_SQL.GetDepartment()
    if err != nil {
        log.Printf("\033[31mError getting departments: %v\033[0m", err)
        http.Error(w, "Internal error, could not retrieve departments.", http.StatusInternalServerError)
        return
    }

	data := struct {
		AllEmployees []Projet_Final_SQL.CompletEmployee
		Department  []Projet_Final_SQL.Department

	}{
		AllEmployees: allEmployeesList,
		Department:  department,

	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
