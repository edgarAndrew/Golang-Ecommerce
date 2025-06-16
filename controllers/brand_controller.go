package controllers

import (
	"Golang-Rest-API/database"
	"Golang-Rest-API/errors"
	"Golang-Rest-API/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

// CreateBrand - Admin-only
func CreateBrand(c *fiber.Ctx) error {
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

	brand := new(models.Brand)
	if err := c.BodyParser(brand); err != nil {
		log.Print(err, "Invalid request body")
		return errors.ErrBadRequest
	}

	if err := database.DB.Create(&brand).Error; err != nil {
		log.Print(err, "Failed to create brand")
		return errors.ErrInternalServer
	}

	return c.Status(fiber.StatusCreated).JSON(brand)
}

// GetBrand - Fetch a single brand by ID
func GetBrand(c *fiber.Ctx) error {
	brandID := c.Params("id")
	var brand models.Brand

	if err := database.DB.First(&brand, brandID).Error; err != nil {
		log.Print(err, "Brand not found")
		return errors.ErrNotFound
	}

	return c.JSON(brand)
}

// GetAllBrands - Fetch all brands
func GetAllBrands(c *fiber.Ctx) error {
	var brands []models.Brand

	if err := database.DB.Find(&brands).Error; err != nil {
		log.Print(err, "Failed to fetch brands")
		return errors.ErrInternalServer
	}

	return c.JSON(fiber.Map{"brands": brands})
}

// DeleteBrand - Admin-only, deletes a brand by ID
func DeleteBrand(c *fiber.Ctx) error {
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

	brandID := c.Params("id")
	var brand models.Brand

	if err := database.DB.First(&brand, brandID).Error; err != nil {
		log.Print(err, "Brand not found")
		return errors.ErrNotFound
	}

	if err := database.DB.Delete(&brand).Error; err != nil {
		log.Print(err, "Failed to delete brand")
		return errors.ErrInternalServer
	}

	return c.JSON(fiber.Map{"message": "Brand deleted successfully"})
}
