package Projet_Final_SQL

import (
	"database/sql"
	"net/http"
)

func GetAllEmployees(w http.ResponseWriter, r *http.Request) []CompletEmployee {

	//Open the database connection
	db, err := sql.Open("sqlite3", "bdd.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, _ := db.Query(`SELECT DISTINCT e.Id, e.LastName, e.FirstName, e.Birthday, e.Phone, e.Address, p.Title AS PostTitle, d.Title AS Department, s.Wage, CONCAT(rf.FirstName, ' ', rf.LastName) AS ReferentName 
	FROM employees e 
	LEFT JOIN posts p ON e.PostId = p.Id 
	LEFT JOIN department d ON p.DepartmentId = d.Id 
	LEFT JOIN salary s ON p.SalaryId = s.Id 
	LEFT JOIN referent r ON e.referentId = r.ReferentId 
	LEFT JOIN employees rf ON r.ReferentId = rf.Id
	`)
	defer rows.Close()

	employeesList := make([]CompletEmployee, 0)

	for rows.Next() {
		employees := CompletEmployee{}
		err = rows.Scan(&employees.Id, &employees.LastName, &employees.FirstName, &employees.Birthday, &employees.Phone, &employees.Address, &employees.PostTitle, &employees.Department, &employees.Wage, &employees.ReferentName)
		CheckErr(err, w, r)
		Nemployees := CompletEmployee{LastName: employees.LastName, FirstName: employees.FirstName, Birthday: employees.Birthday, Phone: employees.Phone, Address: employees.Address, PostTitle: employees.PostTitle, Department: employees.Department, Wage: employees.Wage, ReferentName: employees.ReferentName}

		employeesList = append(employeesList, Nemployees)
	}

	return employeesList
}
