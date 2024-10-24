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

	var filtredEmployeesList []Projet_Final_SQL.CompletEmployee
	filtredEmployeesList = Projet_Final_SQL.GetAllEmployees(w, r)

	filter := r.FormValue("filter")

	IdDeletedEmploye := r.FormValue("delete")

	if IdDeletedEmploye != "" {
		Projet_Final_SQL.DeleteEmployee(IdDeletedEmploye, w, r)
	}

	switch filter {
	case "default":
		filtredEmployeesList = Projet_Final_SQL.GetAllEmployees(w, r)
	case "wageASC":
		filtredEmployeesList = Projet_Final_SQL.FilterByWageASC(w, r)
	case "wageDESC":
		filtredEmployeesList = Projet_Final_SQL.FilterByWageDESC(w, r)
	case "alphabetASC":
		filtredEmployeesList = Projet_Final_SQL.FilterByAlphabetASC(w, r)
	case "alphabetDESC":
		filtredEmployeesList = Projet_Final_SQL.FilterByAlphabetDESC(w, r)
	case "birthdayASC":
		filtredEmployeesList = Projet_Final_SQL.FilterByBirthdayASC(w, r)
	case "birthdayDESC":
		filtredEmployeesList = Projet_Final_SQL.FilterByBirthdayDESC(w, r)
	}

	department, err := Projet_Final_SQL.GetDepartment()
	if err != nil {
		log.Printf("\033[31mError getting departments: %v\033[0m", err)
		http.Error(w, "Internal error, could not retrieve departments.", http.StatusInternalServerError)
		return
	}

	data := struct {
		AllEmployees []Projet_Final_SQL.CompletEmployee
		Department   []Projet_Final_SQL.Department
	}{
		AllEmployees: filtredEmployeesList,
		Department:   department,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
