package main

import (
	"go-crud/config"
	"go-crud/controllers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// main is the main entry point of the application
func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()

	router := gin.Default()

	controllers.SetUserRoutes(router)

	port := config.Routes()

	router.Run(":" + port)
}
