package main

import (
	"go-crud/config"
	"go-crud/controllers"
	"log"
	"os"

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

	router.POST("/users/create", controllers.CreateUser)
	router.GET("/users/getBy/:id", controllers.GetUserByID)
	router.PUT("/users/update/:id", controllers.UpdateUser)
	router.DELETE("/users/delete/:id", controllers.DeleteUser)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
