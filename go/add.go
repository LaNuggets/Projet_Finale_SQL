package Projet_Final_SQL

import (
	"database/sql"
	"fmt"
	"net/http"
)

type Employees struct {
	Id        int
	LastName  string
	FirstName string
	Birthday  string
	Phone     string
	Address   string
}

func CheckErr(err error, w http.ResponseWriter, r *http.Request) {
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
	}
}

func AddEmployee(employe Employees, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "bdd.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, _ := db.Prepare("INSERT INTO employees (LastName, FirstName, BirthDay, Phone, Address) VALUES (?, ?, ?, ?, ?)")
	query.Exec(employe.LastName, employe.FirstName, employe.Birthday, employe.Phone, employe.Address)
	defer query.Close()
}
