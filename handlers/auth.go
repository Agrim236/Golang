package handlers

import (
	"github.com/gofiber/fiber/v2"
	"golang-notes-api/models"
	"golang-notes-api/utils"
)

func Register(c *fiber.Ctx) error {
	var input models.RegisterInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Bad request"})
	}

	// Check if user exists
	var existingUser models.User
	if err := utils.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		return c.Status(409).JSON(fiber.Map{"error": "Email already exists"})
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not hash password"})
	}

	// Create user
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
	}

	if err := utils.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not create user"})
	}

	return c.Status(201).JSON(fiber.Map{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}

func Login(c *fiber.Ctx) error {
	var input models.LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Bad request"})
	}

	// Find user
	var user models.User
	if err := utils.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Check password
	if !utils.CheckPasswordHash(input.Password, user.Password) {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Generate JWT
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not generate token"})
	}

	return c.JSON(fiber.Map{"token": token})
}
