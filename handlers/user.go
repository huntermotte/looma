package handlers

import (
    "context"
    "net/http"
    "strconv"
    "time"
    "gin/models"
    "gin/utils"
    "github.com/gin-gonic/gin"
)

func GetUserRecentTasksAndNotes(c *gin.Context) {
    // Create a context with cancellation to handle client disconnects
    ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
    defer cancel()

    // Get user_id from the URL parameters
    userID, err := strconv.Atoi(c.Param("user_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    // Get limit query parameter (default to 10 if not provided)
    limitParam := c.DefaultQuery("limit", "10")
    limit, err := strconv.Atoi(limitParam)
    if err != nil || limit < 1 {
        limit = 10
    }

    // Fetch user information
    user, err := models.GetUserByID(userID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    // Channels to receive results from concurrent goroutines
    notesChan := make(chan []utils.Note, 1)
    tasksChan := make(chan []utils.Task, 1)
    errChan := make(chan error, 2)

    // Start fetching tasks and notes concurrently using goroutines
    go func() {
        notes, err := utils.ReadRecentNotes(ctx, userID, limit)
        if err != nil {
            errChan <- err
        } else {
            notesChan <- notes
        }
    }()

    go func() {
        tasks, err := utils.FetchRecentTasks(ctx, userID, limit)
        if err != nil {
            errChan <- err
        } else {
            tasksChan <- tasks
        }
    }()

    // Wait for tasks and notes, or handle errors
    var notes []utils.Note
    var tasks []utils.Task

    for i := 0; i < 2; i++ {
        select {
        case <-ctx.Done():
            c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request canceled by client"})
            return
        case err := <-errChan:
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data", "details": err.Error()})
            return
        case fetchedNotes := <-notesChan:
            notes = fetchedNotes
        case fetchedTasks := <-tasksChan:
            tasks = fetchedTasks
        }
    }

    // Return the combined result
    c.JSON(http.StatusOK, gin.H{
        "user_id":   user.ID,
        "user_name": user.Name,
        "tasks":     tasks,
        "notes":     notes,
    })
}
