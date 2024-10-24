package Projet_Final_SQL

import (
	"database/sql"
)

type Department struct {
    Id   int
    Title string
}

func GetDepartment() ([]Department, error) {
    db, err := sql.Open("sqlite3", "bdd.db?_foreign_keys=on")
    if err != nil {
        return nil, err
    }
    defer db.Close()

    rows, err := db.Query("SELECT Id, Title FROM department")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var department []Department
    for rows.Next() {
        var depart Department
        if err := rows.Scan(&depart.Id, &depart.Title); err != nil {
            return nil, err
        }
        department = append(department, depart)
    }
    return department, nil
}