package main

import (
	"fmt"
	"log"

	"golang-notes-api/models" 
	"golang-notes-api/utils"
)

func main() {
	// Connect to database
	err := utils.ConnectDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	
	// Auto-migrate the models
	err = utils.DB.AutoMigrate(&models.User{}, &models.Note{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
		fmt.Println("✅ Database connection successful and migrations completed")
	// First, try to find an existing user with this email (including soft-deleted)
	var existingUser models.User
	result := utils.DB.Unscoped().Where("email = ?", "test@example.com").First(&existingUser)
		// If the user exists, delete it and its associated notes
	if result.Error == nil {
		// Delete associated notes first (to maintain referential integrity)
		utils.DB.Unscoped().Where("user_id = ?", existingUser.ID).Delete(&models.Note{})
		// Then delete the user with hard delete (Unscoped)
		utils.DB.Unscoped().Delete(&existingUser)
		fmt.Println("🗑️ Deleted existing test user and associated notes")
	} else {
		fmt.Println("ℹ️ No existing test user found, will create a new one")
	}
	
	// Create a test user
	testUser := models.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123", // In a real app, this would be hashed
	}
		// Try to save to the database
	result = utils.DB.Create(&testUser)
	if result.Error != nil {
		log.Fatalf("Failed to create test user: %v", result.Error)
	}
	
	fmt.Printf("✅ Test user created with ID: %d\n", testUser.ID)
	
	// Create a test note linked to the user
	testNote := models.Note{
		Title:   "Test Note",
		Content: "This is a test note",
		UserID:  testUser.ID,
	}
		// Save the note
	result = utils.DB.Create(&testNote)
	if result.Error != nil {
		log.Fatalf("Failed to create test note: %v", result.Error)
	}
	
	fmt.Printf("✅ Test note created with ID: %d\n", testNote.ID)
	
	// Retrieve the user with their notes
	var retrievedUser models.User
	result = utils.DB.Preload("Notes").First(&retrievedUser, testUser.ID)
	if result.Error != nil {
		log.Fatalf("Failed to retrieve user with notes: %v", result.Error)
	}
	
	fmt.Printf("✅ Retrieved user with %d notes\n", len(retrievedUser.Notes))
	
	fmt.Println("All tests passed successfully!")
}
