package Projet_Final_SQL

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

func EditEmployee(w http.ResponseWriter, r *http.Request) {
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
	postTitle := r.FormValue("post")
	ReferentName := r.FormValue("manageBy")

	var postId int
	var referentId int

	fmt.Println(lastName, firstName, phone, address, birthday, department, postTitle, ReferentName)

	db, err := sql.Open("sqlite3", "bdd.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	err = db.QueryRow("SELECT Id FROM posts WHERE Title = ?", postTitle).Scan(&postId)
	if err != nil {
		fmt.Println("Aucun post trouvé avec ce titre.")
		return
	}

	err = db.QueryRow("SELECT Id FROM employees WHERE LOWER(FirstName) = ?", strings.ToLower(ReferentName)).Scan(&referentId)
	if err != nil {
		fmt.Println("Aucun manager trouvé avec ce prénom.")
		return
	}
	newEmployee := SendCompletEmployee{LastName: lastName, FirstName: firstName, Phone: phone, Address: address, Birthday: birthday, PostId: postId, ReferentId: referentId}

	query, _ := db.Prepare("UPDATE employees SET (LastName, FirstName, BirthDay, Phone, Address) VALUES (?, ?, ?, ?, ?, ?, ?)")
	query.Exec(newEmployee.LastName, newEmployee.FirstName, newEmployee.Birthday, newEmployee.Phone, newEmployee.Address, newEmployee.PostId, newEmployee.ReferentId)
	defer query.Close()

	http.Redirect(w, r, "/allemployees", http.StatusSeeOther)
}
