package main

import (
	"GoProductManager/config"
	"GoProductManager/database"
	"GoProductManager/middlewares"
	"GoProductManager/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()        // Load configuration
	database.ConnectDatabase() // Connect to PostgreSQL
	router := gin.Default()    // Initialize router
	router.Use(middlewares.LoggingMiddleware())
	routes.RegisterRoutes(router)

	log.Println("Server is running on port 8080")
	router.Run(":8080")
}
