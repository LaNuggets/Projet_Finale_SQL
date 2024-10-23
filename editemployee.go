package main

import (
	Projet_Final_SQL "Projet_Final_SQL/go"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func EditEmployeePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		Projet_Final_SQL.EditEmployee(w, r)
		return
	}
	fmt.Println("non")

	tmpl, err := template.ParseFiles("./static/html/editemployee.html") // Read the home page
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
