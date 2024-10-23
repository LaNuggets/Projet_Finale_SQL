package Projet_Final_SQL

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

func EditEmployee(EditEmployeeFirstName string, w http.ResponseWriter, r *http.Request) {
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

	postTitle := r.FormValue("post")
	ReferentName := r.FormValue("manageBy")

	var postId int
	var referentId int

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

	query, _ := db.Prepare("UPDATE employees SET LastName = ?, FirstName = ?, BirthDay = ?, Phone = ?, Address = ?, PostId = ?, referentId = ? WHERE LOWER(FirstName) = ?")
	query.Exec(newEmployee.LastName, newEmployee.FirstName, newEmployee.Birthday, newEmployee.Phone, newEmployee.Address, newEmployee.PostId, newEmployee.ReferentId, strings.ToLower(EditEmployeeFirstName))
	defer query.Close()

	http.Redirect(w, r, "/allemployees", http.StatusSeeOther)
}
