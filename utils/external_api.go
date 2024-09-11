package utils

import (
    "context"
    "time"
)

type Task struct {
    Timestamp time.Time `json:"timestamp"`
    UserID    int       `json:"user_id"`
    Task      string    `json:"task"`
}

// Mock function that returns tasks for a specific user in sorted order
func FetchRecentTasks(ctx context.Context, userID int, limit int) ([]Task, error) {
    // Mock tasks with different timestamps and user IDs
    mockTasks := []Task{
        {Timestamp: time.UnixMicro(1694083200000000), UserID: 1, Task: "Submit code review"},
        {Timestamp: time.UnixMicro(1694169600000000), UserID: 2, Task: "Update project timeline"},
        {Timestamp: time.UnixMicro(1694256000000000), UserID: 3, Task: "Fix login bug"},
        {Timestamp: time.UnixMicro(1694342400000000), UserID: 1, Task: "Prepare presentation"},
        {Timestamp: time.UnixMicro(1694428800000000), UserID: 4, Task: "Deploy new version"},
        {Timestamp: time.UnixMicro(1694515200000000), UserID: 2, Task: "Write test cases"},
        {Timestamp: time.UnixMicro(1694601600000000), UserID: 3, Task: "Optimize database queries"},
        {Timestamp: time.UnixMicro(1694688000000000), UserID: 4, Task: "Document new feature"},
        {Timestamp: time.UnixMicro(1694774400000000), UserID: 1, Task: "Review team feedback"},
        {Timestamp: time.UnixMicro(1694860800000000), UserID: 2, Task: "Update API documentation"},
    }

    var userTasks []Task
    for _, task := range mockTasks {
        select {
        case <-ctx.Done():
            return nil, ctx.Err() // Handle context cancellation
        default:
        }

        if task.UserID == userID {
            userTasks = append(userTasks, task)
        }
    }

    // Sort tasks by timestamp in ascending order (already sorted in this case)
    if len(userTasks) > limit {
        return userTasks[len(userTasks)-limit:], nil
    }
    return userTasks, nil
}
