package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Ouvrir la base de données (ou la créer si elle n'existe pas)
	db, err := sql.Open("sqlite3", "../bdd.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Lire le fichier SQL
	sqlFile, err := ioutil.ReadFile("../insert.sql")
	if err != nil {
		log.Fatal(err)
	}

	// Exécuter les commandes SQL
	_, err = db.Exec(string(sqlFile))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Insertion réussie !")
}
