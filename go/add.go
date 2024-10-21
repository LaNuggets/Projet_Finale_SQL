package Projet_Final_SQL

import (
	"database/sql"
	"fmt"
	"net/http"
)

type Employee struct {
	Id         int
	LastName   string
	FirstName  string
	Birthday   string
	Phone      string
	Address    string
	PostId     int
	ReferentId int
}

func CheckErr(err error, w http.ResponseWriter, r *http.Request) {
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
	}
}

func AddEmployee(w http.ResponseWriter, r *http.Request) {
	errP := r.ParseForm()
	if errP != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	lastName := r.FormValue("lastName")
	firstName := r.FormValue("firstName")
	phone := r.FormValue("phoneNumber")
	address := r.FormValue("adress")
	birthday := r.FormValue("birthday")

	department := r.FormValue("department")
	post := r.FormValue("post")
	manageBy := r.FormValue("manageBy")

	fmt.Println(lastName, firstName, phone, address, birthday, department, post, manageBy)

	newEmployee := Employee{LastName: lastName, FirstName: firstName, Phone: phone, Address: address, Birthday: birthday}

	//Open the database connection
	db, err := sql.Open("sqlite3", "bdd.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, _ := db.Prepare("INSERT INTO employees (LastName, FirstName, BirthDay, Phone, Address) VALUES (?, ?, ?, ?, ?)")
	query.Exec(newEmployee.LastName, newEmployee.FirstName, newEmployee.Birthday, newEmployee.Phone, newEmployee.Address)
	defer query.Close()

	http.Redirect(w, r, "/allemployees", http.StatusSeeOther)
}
