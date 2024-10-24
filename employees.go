package main

import (
	Projet_Final_SQL "Projet_Final_SQL/go"
	"html/template"
	"log"
	"net/http"
)

func Employees(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		Projet_Final_SQL.AddEmployee(w, r)
		return
	}

	tmpl, err := template.ParseFiles("./static/html/employees.html") // Read the home page
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	department, err := Projet_Final_SQL.GetDepartment()
    if err != nil {
        log.Printf("\033[31mError getting all departments: %v\033[0m", err)
        http.Error(w, "Internal error, could not retrieve departments.", http.StatusInternalServerError)
        return
    }

	referents, err := Projet_Final_SQL.GetReferents()
    if err != nil {
        log.Printf("\033[31mError getting all referents: %v\033[0m", err)
        http.Error(w, "Internal error, could not retrieve referents.", http.StatusInternalServerError)
        return
    }
	
	posts, err := Projet_Final_SQL.GetPostsByDepartment()
    if err != nil {
        log.Printf("\033[31mError getting all referents: %v\033[0m", err)
        http.Error(w, "Internal error, could not retrieve referents.", http.StatusInternalServerError)
        return
    }

    data := struct {
        Departments []Projet_Final_SQL.Department
		Referents   []Projet_Final_SQL.Referent
		Posts       []Projet_Final_SQL.Post

    }{
        Departments: department,
		Referents:   referents,
		Posts:       posts,
    }

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
