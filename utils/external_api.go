package utils

import (
    "strconv"
    "fmt"
    "context"
    "math/rand"
    "time"
)

type Task struct {
    Timestamp time.Time `json:"timestamp"`
    UserID    int       `json:"user_id"`
    Task      string    `json:"task"`
}

// In-memory task storage (for testing purposes)
var Tasks []Task

// GenerateTasks generates a specified number of tasks for random users.
func GenerateTasks(numTasks int, numUsers int) {
    // Clear existing tasks
    Tasks = make([]Task, 0, numTasks)

    rand.Seed(time.Now().UnixNano())

    for i := 0; i < numTasks; i++ {
        // Random user ID between 1 and numUsers
        userID := rand.Intn(numUsers) + 1

        // Random timestamp in microseconds
        timestamp := time.Now().Add(time.Duration(i) * time.Second).UnixMicro()

        // Create a new task with a simple description
        task := Task{
            Timestamp: time.UnixMicro(timestamp),
            UserID:    userID,
            Task:      "Task " + strconv.Itoa(i+1),
        }

        // Add the task to the task list
        Tasks = append(Tasks, task)
    }

    // Optionally log or print the number of tasks generated
    fmt.Printf("Generated %d tasks for %d users\n", numTasks, numUsers)
}

// FetchRecentTasks returns recent tasks for the specified user from the in-memory task list.
func FetchRecentTasks(ctx context.Context, userID int, limit int) ([]Task, error) {
    var userTasks []Task

    // Iterate through the in-memory list of tasks and find the tasks for the user
    for _, task := range Tasks {
        if task.UserID == userID {
            userTasks = append(userTasks, task)

            // Stop if we've reached the limit
            if len(userTasks) == limit {
                break
            }
        }
    }

    return userTasks, nil
}
