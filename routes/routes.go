package routes

import (
    "github.com/gofiber/fiber/v2"
    "golang-notes-api/handlers"
    "golang-notes-api/middleware"
)

// SetupRoutes configures all the routes for our application
func SetupRoutes(app *fiber.App) {
    // Public routes
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Notes API is running")
    })

    // Auth routes (public)
    app.Post("/register", handlers.Register)
    app.Post("/login", handlers.Login)
    
    // Notes routes (protected by JWT middleware)
    notes := app.Group("/notes")
    notes.Use(middleware.AuthMiddleware)
    
    notes.Post("/", handlers.CreateNote)
    notes.Get("/", handlers.GetNotes)
    notes.Get("/:id", handlers.GetNote)
    notes.Put("/:id", handlers.UpdateNote)
    notes.Delete("/:id", handlers.DeleteNote)
}