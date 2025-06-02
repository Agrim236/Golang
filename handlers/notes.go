package handlers

import (
    "github.com/gofiber/fiber/v2"
    "golang-notes-api/models"
    "golang-notes-api/utils"
    "strconv"
)

// CreateNote - Handler for creating a new note
func CreateNote(c *fiber.Ctx) error {
    // 1. Get authenticated user from JWT middleware
    userID := c.Locals("userID").(uint)

    // 2. Parse request body
    var noteInput struct {
        Title   string `json:"title" validate:"required"`
        Content string `json:"content" validate:"required"`
    }
    
    if err := c.BodyParser(&noteInput); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    // 3. Validate input
    if noteInput.Title == "" || noteInput.Content == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Title and content are required",
        })
    }

    // 4. Create note in database
    note := models.Note{
        Title:   noteInput.Title,
        Content: noteInput.Content,
        UserID:  userID, // Automatically associate with logged-in user
    }

    if err := utils.DB.Create(&note).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create note",
        })
    }

    // 5. Return created note
    return c.Status(fiber.StatusCreated).JSON(note)
}

// GetNotes - Handler for fetching all notes of the logged-in user
func GetNotes(c *fiber.Ctx) error {
    // 1. Get authenticated user
    userID := c.Locals("userID").(uint)

    // 2. Fetch pagination parameters (Bonus)
    page, _ := strconv.Atoi(c.Query("page", "1"))
    limit, _ := strconv.Atoi(c.Query("limit", "10"))
    offset := (page - 1) * limit

    // 3. Fetch notes from database
    var notes []models.Note
    query := utils.DB.Where("user_id = ?", userID)
    
    // Bonus: Search functionality
    if search := c.Query("search"); search != "" {
        query = query.Where("title LIKE ? OR content LIKE ?", "%"+search+"%", "%"+search+"%")
    }

    // Bonus: Pagination
    if err := query.Offset(offset).Limit(limit).Find(&notes).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to fetch notes",
        })
    }

    // 4. Return notes list
    return c.JSON(fiber.Map{
        "data":  notes,
        "page":  page,
        "limit": limit,
    })
}

// GetNote - Handler for fetching a single note
func GetNote(c *fiber.Ctx) error {
    // 1. Get authenticated user
    userID := c.Locals("userID").(uint)
    
    // 2. Get note ID from URL params
    noteID, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid note ID",
        })
    }

    // 3. Fetch note from database
    var note models.Note
    if err := utils.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Note not found",
        })
    }

    // 4. Return note
    return c.JSON(note)
}

// UpdateNote - Handler for updating a note
func UpdateNote(c *fiber.Ctx) error {
    // 1. Get authenticated user
    userID := c.Locals("userID").(uint)
    
    // 2. Get note ID from URL params
    noteID, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid note ID",
        })
    }

    // 3. Parse request body
    var noteInput struct {
        Title   string `json:"title" validate:"required"`
        Content string `json:"content" validate:"required"`
    }
    
    if err := c.BodyParser(&noteInput); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    // 4. Validate input
    if noteInput.Title == "" || noteInput.Content == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Title and content are required",
        })
    }

    // 5. Check if note exists and belongs to user
    var note models.Note
    if err := utils.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Note not found",
        })
    }

    // 6. Update note
    note.Title = noteInput.Title
    note.Content = noteInput.Content

    if err := utils.DB.Save(&note).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to update note",
        })
    }

    // 7. Return updated note
    return c.JSON(note)
}

// DeleteNote - Handler for deleting a note
func DeleteNote(c *fiber.Ctx) error {
    // 1. Get authenticated user
    userID := c.Locals("userID").(uint)
    
    // 2. Get note ID from URL params
    noteID, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid note ID",
        })
    }

    // 3. Check if note exists and belongs to user
    var note models.Note
    if err := utils.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Note not found",
        })
    }

    // 4. Delete note
    if err := utils.DB.Delete(&note).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to delete note",
        })
    }

    // 5. Return success message
    return c.JSON(fiber.Map{
        "message": "Note deleted successfully",
    })
}
