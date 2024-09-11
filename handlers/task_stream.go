package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "gin/models"
    "gin/utils"
    "github.com/gin-gonic/gin"
)

// StreamTasks streams all tasks in ascending order of timestamp.
func StreamTasks(c *gin.Context) {
    // Open a database connection
    rows, err := models.DB.Query(`
        SELECT timestamp, user_id, task
        FROM tasks
        ORDER BY timestamp ASC
    `)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query tasks"})
        return
    }
    defer rows.Close()

    // Set the content type to stream JSON
    c.Header("Content-Type", "application/json")

    // Start streaming the tasks in a sorted manner
    enc := json.NewEncoder(c.Writer)
    c.Writer.WriteHeader(http.StatusOK)

    for rows.Next() {
        // Check if the client has disconnected
        select {
        case <-c.Request.Context().Done():
            log.Println("Client disconnected, stopping stream")
            return
        default:
        }

        var task utils.Task
        var timestamp int64

        if err := rows.Scan(&timestamp, &task.UserID, &task.Task); err != nil {
            log.Println("Error scanning row:", err)
            continue
        }

        task.Timestamp = timestamp

        // Encode the task and stream it as JSON
        if err := enc.Encode(task); err != nil {
            log.Println("Error encoding task:", err)
            break
        }

        // Flush the response buffer to keep the stream alive
        c.Writer.Flush()
    }
}
