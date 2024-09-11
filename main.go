package main

import (
	"time"
    "gin/models"
    "gin/handlers"
    "gin/utils"
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    numUsers := 10 // TODO: make env variable?
    numNotes := 500
    numTasks := 500

    start := time.Now()
    models.InitDB()
    models.GenerateUsers(numUsers) // Insert sample users, using configurable number
    utils.CreateNotesFile(numNotes, numUsers)
    utils.GenerateTasks(numTasks, numUsers)

    elapsed := time.Since(start)
    log.Println("Data setup complete in", elapsed)

    router := gin.Default()

    router.GET("/user/:user_id/recent", handlers.GetUserRecentTasksAndNotes)

    router.Run(":8080")
}
