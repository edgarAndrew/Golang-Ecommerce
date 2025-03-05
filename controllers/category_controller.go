package controllers

import (
	"Golang-Rest-API/database"
	"Golang-Rest-API/errors"
	"Golang-Rest-API/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

// CreateCategory - Admin-only
func CreateCategory(c *fiber.Ctx) error {
	// Get user ID from context
	userID := c.Locals("userID")
	var user models.User
	if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
		log.Print(err, "Failed to find user")
		return errors.ErrNotFound
	}

	// Check admin permissions
	if !user.IsAdmin() {
		log.Print("User is not an admin")
		return errors.ErrForbidden
	}

	category := new(models.Category)
	if err := c.BodyParser(category); err != nil {
		log.Print(err, "Invalid request body")
		return errors.ErrBadRequest
	}

	if err := database.DB.Create(&category).Error; err != nil {
		log.Print(err, "Failed to create category")
		return errors.ErrInternalServer
	}

	return c.Status(fiber.StatusCreated).JSON(category)
}

// GetCategories - Fetch all categories
func GetCategories(c *fiber.Ctx) error {
	var categories []models.Category

	if err := database.DB.Find(&categories).Error; err != nil {
		log.Print(err, "Failed to fetch categories")
		return errors.ErrInternalServer
	}

	return c.JSON(fiber.Map{"categories": categories})
}

// DeleteCategory - Deletes a category (only if the user is an admin)
func DeleteCategory(c *fiber.Ctx) error {
	// Get user ID from context
	userID := c.Locals("userID")
	var user models.User
	if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
		log.Print(err, "Failed to find user")
		return errors.ErrNotFound
	}

	// Check admin permissions
	if !user.IsAdmin() {
		log.Print("User is not an admin")
		return errors.ErrForbidden
	}

	categoryID := c.Params("id")
	var category models.Category

	// Check if category exists
	if err := database.DB.First(&category, categoryID).Error; err != nil {
		log.Print(err, "Category not found")
		return errors.ErrNotFound
	}

	// Delete the category
	if err := database.DB.Delete(&category).Error; err != nil {
		log.Print(err, "Failed to delete category")
		return errors.ErrInternalServer
	}

	return c.JSON(fiber.Map{"message": "Category deleted successfully"})
}
