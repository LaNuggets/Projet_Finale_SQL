package Projet_Final_SQL

import (
    "database/sql"
)

type Post struct {
    ID           int
    Title        string
}

func GetPostsByDepartment() ([]Post, error) {
    db, err := sql.Open("sqlite3", "bdd.db?_foreign_keys=on")
    if err != nil {
        return nil, err
    }
    defer db.Close()

    rows, err := db.Query("SELECT Id, Title FROM posts")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var posts []Post
    for rows.Next() {
        var post Post
        if err := rows.Scan(&post.ID, &post.Title); err != nil {
            return nil, err
        }
        posts = append(posts, post)
    }
    return posts, nil
}