package Projet_Final_SQL

import (
    "database/sql"
)

type Referent struct {
    ID   int
    Name string
}

func GetReferents() ([]Referent, error) {
    db, err := sql.Open("sqlite3", "bdd.db?_foreign_keys=on")
    if err != nil {
        return nil, err
    }
    defer db.Close()

    rows, err := db.Query(`
	SELECT DISTINCT e.ReferentId, r.FirstName || ' ' || r.LastName AS Name
	FROM employees e
	JOIN employees r ON e.ReferentId = r.Id
	WHERE e.ReferentId IS NOT NULL
	`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var referents []Referent
    for rows.Next() {
        var referent Referent
        if err := rows.Scan(&referent.ID, &referent.Name); err != nil {
            return nil, err
        }
        referents = append(referents, referent)
    }
    return referents, nil
}