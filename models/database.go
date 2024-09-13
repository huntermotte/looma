package models

import (
    "database/sql"
    "fmt"
    "log"
    "math/rand"
    "time"

    _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("sqlite3", "./users.db")
    if err != nil {
        log.Fatal(err)
    }

    // Create the users table if it doesn't exist
    createUsersTable := `
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL
        );
    `
    _, err = DB.Exec(createUsersTable)
    if err != nil {
        log.Fatal(err)
    }

    // Create the tasks table if it doesn't exist
    createTasksTable := `
        CREATE TABLE IF NOT EXISTS tasks (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL,
            task TEXT NOT NULL,
            timestamp INTEGER NOT NULL
        );
    `
    _, err = DB.Exec(createTasksTable)
    if err != nil {
        log.Fatal(err)
    }

    // Clear out old data (delete all records from users and tasks tables)
    ClearOldData()
}

// ClearOldData removes all records from the users and tasks tables and resets the auto-increment counter
func ClearOldData() {
    // Clear users table
    _, err := DB.Exec("DELETE FROM users")
    if err != nil {
        log.Fatal("Failed to clear users table:", err)
    }

    // Clear tasks table
    _, err = DB.Exec("DELETE FROM tasks")
    if err != nil {
        log.Fatal("Failed to clear tasks table:", err)
    }

    // Reset the auto-increment counter for users and tasks
    _, err = DB.Exec("DELETE FROM sqlite_sequence WHERE name = 'users'")
    if err != nil {
        log.Fatal("Failed to reset user auto-increment counter:", err)
    }

    _, err = DB.Exec("DELETE FROM sqlite_sequence WHERE name = 'tasks'")
    if err != nil {
        log.Fatal("Failed to reset task auto-increment counter:", err)
    }

    fmt.Println("Old data cleared and auto-increment counters reset for users and tasks")
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

// GenerateTasks generates tasks for random users and stores them in the database.
func GenerateTasks(numTasks int, numUsers int) {
    stmt, err := DB.Prepare("INSERT INTO tasks (user_id, task, timestamp) VALUES (?, ?, ?)")
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()

    rand.Seed(time.Now().UnixNano())

    for i := 0; i < numTasks; i++ {
        // Random user ID between 1 and numUsers
        userID := rand.Intn(numUsers) + 1

        // Random timestamp in microseconds
        timestamp := time.Now().Add(time.Duration(i) * time.Second).UnixMicro()

        // Generate the task description
        taskDescription := fmt.Sprintf("Task %d", i+1)

        // Insert task into the database
        _, err = stmt.Exec(userID, taskDescription, timestamp)
        if err != nil {
            log.Fatal(err)
        }
    }

    fmt.Printf("Successfully inserted %d tasks into the database\n", numTasks)
}
