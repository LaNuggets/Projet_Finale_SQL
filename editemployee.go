package main

import (
	Projet_Final_SQL "Projet_Final_SQL/go"
	"html/template"
	"log"
	"net/http"
)

func EditEmployeePage(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("firstname")
	firstName2 := r.FormValue("firstname")

	if r.Method == http.MethodPost {
		Projet_Final_SQL.EditEmployee(firstName2, w, r)
		return
	}

	tmpl, err := template.ParseFiles("./static/html/editemployee.html") // Read the home page
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	data := struct {
		FirstNameSend string
	}{
		FirstNameSend: firstName,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
