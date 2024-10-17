package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var port = ":8766"

func main() {

	css := http.FileServer(http.Dir("../style"))             // For add css to the html pages
	http.Handle("/style/", http.StripPrefix("/style/", css)) // For add css to the html pages

	http.HandleFunc("/accueil", func(w http.ResponseWriter, r *http.Request) { // Start the communication whit the html page
		accueil := template.Must(template.ParseFiles("../html/accueil.html"))
		accueil.Execute(w, nil)
	})
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/nom_base_de_donnée") // Open the connection whit the database
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) { // Start the communication whit the html page
		employees := template.Must(template.ParseFiles("../html/employees.html"))
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusInternalServerError)
			return
		}
		// Récupère les données du formulaire
		lastName := r.FormValue("lastName")
		firstName := r.FormValue("firstName")
		phoneNumber := r.FormValue("phoneNumber")
		adress := r.FormValue("adress")
		birthday := r.FormValue("birthday")

		// Insère les données dans la base de données
		_, err = db.Exec("INSERT INTO employees (LastName, FirstName, Birthday, phoneNumber, adress) VALUES (?, ?, ?, ?, ?)", lastName, firstName, birthday, phoneNumber, adress)
		if err != nil {
			http.Error(w, "Error inserting data into database", http.StatusInternalServerError)
			return
		}
		employees.Execute(w, nil)
	})

	fmt.Println("http://localhost" + port + "/accueil")
	http.ListenAndServe(port, nil)
}
