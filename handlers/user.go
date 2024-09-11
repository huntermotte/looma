package handlers

import (
    "net/http"
    "strconv"
    "gin/models"
    "gin/utils"
    "github.com/gin-gonic/gin"
)

func GetUserRecentTasksAndNotes(c *gin.Context) {
    userID, err := strconv.Atoi(c.Param("user_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    user, err := models.GetUserByID(userID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    notes, err := utils.ReadRecentNotes(userID, 10)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read notes"})
        return
    }

    tasks, err := utils.FetchRecentTasks(userID, 10)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "user_id":   user.ID,
        "user_name": user.Name,
        "tasks":     tasks,
        "notes":     notes,
    })
}
