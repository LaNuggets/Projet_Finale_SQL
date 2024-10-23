package main

import (
	Projet_Final_SQL "Projet_Final_SQL/go"

	"html/template"
	"log"
	"net/http"
)

func Accueil(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("./static/html/accueil.html") // Read the home page
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}
	
	departments, err := Projet_Final_SQL.GetDepartment()
    if err != nil {
        log.Printf("\033[31mError getting all departments: %v\033[0m", err)
        http.Error(w, "Internal error, could not retrieve departments.", http.StatusInternalServerError)
        return
    }

    data := struct {
        Departments []Projet_Final_SQL.Department
    }{
        Departments: departments,
    }

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	
}
