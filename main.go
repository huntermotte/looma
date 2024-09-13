package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"gin/handlers"
	"gin/models"
	"gin/utils"
	"github.com/gin-gonic/gin"
)

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	// Convert string to int
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("Invalid value for %s, using default: %d\n", key, defaultValue)
		return defaultValue
	}
	return value
}

func main() {
	// Retrieve environment variables, or use defaults
	numUsers := getEnvAsInt("NUM_USERS", 10)
	numNotes := getEnvAsInt("NUM_NOTES", 500)
	numTasks := getEnvAsInt("NUM_TASKS", 500)

	start := time.Now()

	// Initialize database and generate data
	models.InitDB()
	models.GenerateUsers(numUsers) // Insert sample users, using configurable number
	utils.CreateNotesFile(numNotes, numUsers)
	models.GenerateTasks(numTasks, numUsers)

	elapsed := time.Since(start)
	log.Println("Data setup complete in", elapsed)

	// Initialize the router
	router := gin.Default()

	// API Endpoints
	router.GET("/user/:user_id/recent", handlers.GetUserRecentTasksAndNotes)

	// External API endpoint for streaming tasks
	router.GET("/external/tasks/stream", handlers.StreamTasks)

	// Start the server
	router.Run(":8080")
}
