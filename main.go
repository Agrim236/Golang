package main

import (
	"log"
	"golang-notes-api/routes"
	"golang-notes-api/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Connect to the database using utils package
	err := utils.ConnectDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	
	// Set up routes and start the server
	app := fiber.New()
	routes.SetupRoutes(app)

	err = app.Listen(":3001")  // Changed from 3000 to 3001
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
