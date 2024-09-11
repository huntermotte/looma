package main

import (
    "gin/models"
    "gin/handlers"
    "github.com/gin-gonic/gin"
)

func main() {
    models.InitDB()
    models.AddSampleUsers() // Insert sample users

    router := gin.Default()

    router.GET("/user/:user_id/recent", handlers.GetUserRecentTasksAndNotes)

    router.Run(":8080")
}
