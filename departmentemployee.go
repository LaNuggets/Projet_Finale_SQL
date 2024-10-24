package main

import (
	Projet_Final_SQL "Projet_Final_SQL/go"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func DepartmentEmployee(w http.ResponseWriter, r *http.Request) {
	departmentID := strings.TrimPrefix(r.URL.Path, "/department/")

	tmpl, err := template.ParseFiles("./static/html/department.html")
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	employeesList := Projet_Final_SQL.GetEmployeesByDepartment(w, r, departmentID)

	// filter := r.FormValue("filter")

	IdDeletedEmploye := r.FormValue("delete")

	// switch filter {
	// case "default":
	// 	filtredEmployeesList = Projet_Final_SQL.GetAllEmployees(w, r)
	// case "wageASC":
	// 	filtredEmployeesList = Projet_Final_SQL.FilterByWageASC(w, r)
	// case "wageDESC":
	// 	filtredEmployeesList = Projet_Final_SQL.FilterByWageDESC(w, r)
	// case "alphabetASC":
	// 	filtredEmployeesList = Projet_Final_SQL.FilterByAlphabetASC(w, r)
	// case "alphabetDESC":
	// 	filtredEmployeesList = Projet_Final_SQL.FilterByAlphabetDESC(w, r)
	// case "birthdayASC":
	// 	filtredEmployeesList = Projet_Final_SQL.FilterByBirthdayASC(w, r)
	// case "birthdayDESC":
	// 	filtredEmployeesList = Projet_Final_SQL.FilterByBirthdayDESC(w, r)
	// }

	if IdDeletedEmploye != "" {
		Projet_Final_SQL.DeleteEmployee(IdDeletedEmploye, w, r)
	}

	department, err := Projet_Final_SQL.GetDepartment()
	if err != nil {
		log.Printf("\033[31mError getting departments: %v\033[0m", err)
		http.Error(w, "Internal error, could not retrieve departments.", http.StatusInternalServerError)
		return
	}

	data := struct {
		DepartmentID string
		Employees    []Projet_Final_SQL.CompletEmployee
		Departments  []Projet_Final_SQL.Department
	}{
		DepartmentID: departmentID,
		Employees:    employeesList,
		Departments:  department,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
