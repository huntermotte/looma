package models

import (
    "fmt"
    "math/rand"
    "time"
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

func GenerateUsers(numUsers int) {
    // Prepare a statement for inserting users
    stmt, err := DB.Prepare("INSERT INTO users (name) VALUES (?)")
    if err != nil {
        fmt.Println("Error preparing insert statement:", err)
        return
    }
    defer stmt.Close()

    // Set up random seed
    rand.Seed(time.Now().UnixNano())

    for i := 0; i < numUsers; i++ {
        // Generate a random user name
        userName := fmt.Sprintf("User%d", i+1)

        // Insert the user into the database
        _, err := stmt.Exec(userName)
        if err != nil {
            fmt.Println("Error inserting user:", err)
        }
    }

    fmt.Printf("Generated %d users in the database\n", numUsers)
}
