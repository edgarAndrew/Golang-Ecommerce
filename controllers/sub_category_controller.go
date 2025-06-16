package controllers

import (
	"Golang-Rest-API/database"
	"Golang-Rest-API/errors"
	"Golang-Rest-API/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

// CreateSubCategory - Admin-only
func CreateSubCategory(c *fiber.Ctx) error {
	// Get user ID from context
	userID := c.Locals("userID")
	var user models.User
	if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
		log.Print(err, "Failed to find user")
		return errors.ErrNotFound
	}


	subCategory := new(models.SubCategory)
	if err := c.BodyParser(subCategory); err != nil {
		log.Print(err, "Invalid request body")
		return errors.ErrBadRequest
	}

	if err := database.DB.Create(&subCategory).Error; err != nil {
		log.Print(err, "Failed to create subcategory")
		return errors.ErrInternalServer
	}

	return c.Status(fiber.StatusCreated).JSON(subCategory)
}

// DeleteSubCategory - Delete a subcategory by ID (Admin-only)
func DeleteSubCategory(c *fiber.Ctx) error {
	// Get user ID from context
	userID := c.Locals("userID")
	var user models.User
	if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
		log.Print(err, "Failed to find user")
		return errors.ErrNotFound
	}


	subCategoryID := c.Params("id") // Extract subcategory ID from the URL

	// Find the subcategory to be deleted
	subCategory := new(models.SubCategory)
	if err := database.DB.Where("id = ?", subCategoryID).First(&subCategory).Error; err != nil {
		log.Print(err, "Subcategory not found")
		return errors.ErrNotFound
	}

	// Delete the subcategory
	if err := database.DB.Delete(&subCategory).Error; err != nil {
		log.Print(err, "Failed to delete subcategory")
		return errors.ErrInternalServer
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Subcategory deleted successfully",
	})
}

// GetSubCategoriesByCategoryID - Get all subcategories for a particular category
func GetSubCategoriesByCategoryID(c *fiber.Ctx) error {
	categoryID := c.Params("category_id")
	var subCategories []models.SubCategory
	log.Print(categoryID)

	// Query all subcategories for the given CategoryID
	if err := database.DB.Where("category_id = ?", categoryID).Find(&subCategories).Error; err != nil {
		log.Print(err, "Failed to fetch subcategories")
		return errors.ErrInternalServer
	}

	if len(subCategories) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No subcategories found for this category",
		})
	}

	return c.JSON(fiber.Map{"subcategories": subCategories})
}
