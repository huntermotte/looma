package handlers

import (
    "net/http"
    "strconv"
    "gin/models"
    "gin/utils"
    "github.com/gin-gonic/gin"
)

func GetUserRecentTasksAndNotes(c *gin.Context) {
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

    // Fetch recent notes for the user
    notes, err := utils.ReadRecentNotes(userID, limit)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read notes"})
        return
    }

    // Fetch recent tasks for the user
    tasks, err := utils.FetchRecentTasks(userID, limit)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
        return
    }

    // Return the response
    c.JSON(http.StatusOK, gin.H{
        "user_id":   user.ID,
        "user_name": user.Name,
        "tasks":     tasks,
        "notes":     notes,
    })
}
