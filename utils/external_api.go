package utils

import (
    "fmt"
    "context"
    "encoding/json"
    "net/http"
)

type Task struct {
    Timestamp int64     `json:"timestamp"`
    UserID    int       `json:"user_id"`
    Task      string    `json:"task"`
}

// StreamTasksFromAPI calls the external API to get a stream of tasks and filters by user_id.
func StreamTasksFromAPI(ctx context.Context, userID int, limit int) ([]Task, error) {
    url := "http://localhost:8080/external/tasks/stream"
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return nil, err
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("Failed to fetch tasks, status: %d", resp.StatusCode)
    }

    dec := json.NewDecoder(resp.Body)
    var tasks []Task
    count := 0

    for dec.More() {
        var task Task
        if err := dec.Decode(&task); err != nil {
            return nil, err
        }

        // Filter tasks by userID
        if task.UserID == userID {
            tasks = append(tasks, task)
            count++
            if count >= limit {
                break
            }
        }
    }

    return tasks, nil
}
