package Projet_Final_SQL

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
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

	postTitle := r.FormValue("posts")
	ReferentName := r.FormValue("manageBy")

	fmt.Println(postTitle + "post")
	fmt.Println(ReferentName + "manage")

	postId, _ := strconv.Atoi(postTitle)
	referentId, _ := strconv.Atoi(ReferentName)

	db, err := sql.Open("sqlite3", "bdd.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	newEmployee := SendCompletEmployee{LastName: lastName, FirstName: firstName, Phone: phone, Address: address, Birthday: birthday, PostId: postId, ReferentId: referentId}

	query, _ := db.Prepare("UPDATE employees SET LastName = ?, FirstName = ?, BirthDay = ?, Phone = ?, Address = ?, PostId = ?, referentId = ? WHERE LOWER(FirstName) = ?")
	query.Exec(newEmployee.LastName, newEmployee.FirstName, newEmployee.Birthday, newEmployee.Phone, newEmployee.Address, newEmployee.PostId, newEmployee.ReferentId, strings.ToLower(EditEmployeeFirstName))
	defer query.Close()

	http.Redirect(w, r, "/allemployees", http.StatusSeeOther)
}
