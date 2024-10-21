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

	rows, _ := db.Query("SELECT e.Id, e.LastName, e.FirstName, e.Birthday, e.Phone, e.Address, e.PostId, e.referentId, d.Title, s.Wage FROM employees e LEFT JOIN posts p ON e.PostId = p.Id LEFT JOIN department d ON p.DepartmentId = d.Id LEFT JOIN salary s ON p.SalaryId = s.Id")
	defer rows.Close()

	employeesList := make([]CompletEmployee, 0)

	for rows.Next() {
		employees := CompletEmployee{}
		err = rows.Scan(&employees.Id, &employees.LastName, &employees.FirstName, &employees.Birthday, &employees.Phone, &employees.Address, &employees.PostId, &employees.ReferentId, &employees.Department, &employees.Wage)
		CheckErr(err, w, r)
		Nemployees := CompletEmployee{LastName: employees.LastName, FirstName: employees.FirstName, Birthday: employees.Birthday, Phone: employees.Phone, Address: employees.Address, PostId: 0, ReferentId: 0}

		employeesList = append(employeesList, Nemployees)
	}

	return employeesList
}
