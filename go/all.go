package Projet_Final_SQL

import (
	"database/sql"
	"net/http"
)

func GetAllEmployees(w http.ResponseWriter, r *http.Request) []Employee {

	//Open the database connection
	db, err := sql.Open("sqlite3", "bdd.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM employees")
	defer rows.Close()

	employeesList := make([]Employee, 0)

	for rows.Next() {
		employees := Employee{}
		err = rows.Scan(&employees.Id, &employees.LastName, &employees.FirstName, &employees.Birthday, &employees.Phone, &employees.Address, &employees.PostId, &employees.ReferentId)
		CheckErr(err, w, r)
		Nemployees := Employee{LastName: employees.LastName, FirstName: employees.FirstName, Birthday: employees.Birthday, Phone: employees.Phone, Address: employees.Address, PostId: 0, ReferentId: 0}

		employeesList = append(employeesList, Nemployees)
	}
	http.Redirect(w, r, "/allemployees", http.StatusSeeOther)

	return employeesList
}
