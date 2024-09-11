package models

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("sqlite3", "./users.db")
    if err != nil {
        log.Fatal(err)
    }

    createTable := `
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL
        );
    `
    _, err = DB.Exec(createTable)
    if err != nil {
        log.Fatal(err)
    }
}

func AddSampleUsers() {
    users := []User{
        {Name: "Alice"},
        {Name: "Bob"},
        {Name: "Charlie"},
        {Name: "Diana"},
    }

    for _, user := range users {
        query := `INSERT INTO users (name) VALUES (?)`
        _, err := DB.Exec(query, user.Name)
        if err != nil {
            log.Println("Error inserting user:", err)
        } else {
            log.Println("Inserted user:", user.Name)
        }
    }
}
