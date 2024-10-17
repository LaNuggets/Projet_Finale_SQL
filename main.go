package main

import (
	Projet_Final_SQL "Projet_Final_SQL/go"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var port = ":8766"

func main() {

	css := http.FileServer(http.Dir("./static/style"))       // For add css to the html pages
	http.Handle("/style/", http.StripPrefix("/style/", css)) // For add css to the html pages

	http.HandleFunc("/accueil", func(w http.ResponseWriter, r *http.Request) { // Start the communication whit the html page
		accueil := template.Must(template.ParseFiles("./static/html/accueil.html"))
		accueil.Execute(w, nil)
	})

	db, err := sql.Open("sqlite3", "bdd.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Tester la connexion
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) { // Start the communication whit the html page
		employees := template.Must(template.ParseFiles("./static/html/employees.html"))
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusInternalServerError)
			return
		}

		// Récupère les données du formulaire
		lastName := r.FormValue("lastName")
		firstName := r.FormValue("firstName")
		phone := r.FormValue("phoneNumber")
		address := r.FormValue("adress")
		birthday := r.FormValue("birthday")

		newEmployee := Projet_Final_SQL.Employees{LastName: lastName, FirstName: firstName, Phone: phone, Address: address, Birthday: birthday}

		Projet_Final_SQL.AddEmployee(newEmployee, w, r)

		employees.Execute(w, nil)
	})

	fmt.Println("http://localhost" + port + "/accueil")
	http.ListenAndServe(port, nil)
}
