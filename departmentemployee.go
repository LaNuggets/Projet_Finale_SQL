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


	allEmployeesList := Projet_Final_SQL.GetAllEmployees(w, r)

    employeesList := Projet_Final_SQL.GetEmployeesByDepartment(w, r, departmentID)


	IdDeletedEmploye := r.FormValue("delete")

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
        AllEmployees []Projet_Final_SQL.CompletEmployee



    }{
        DepartmentID: departmentID,
        Employees:    employeesList,
		Departments:  department,
        AllEmployees: allEmployeesList,
    }

    err = tmpl.Execute(w, data)
    if err != nil {
        log.Printf("\033[31mError executing template: %v\033[0m", err)
        http.Error(w, "Internal error", http.StatusInternalServerError)
        return
    }
}
