package controllers

import (
	"Golang-Rest-API/database"
	"Golang-Rest-API/errors"
	"Golang-Rest-API/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

// AddAddress - Adds a new address for the logged-in user
func AddAddress(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)

	address := new(models.Address)
	if err := c.BodyParser(address); err != nil {
		log.Print(err, "Invalid request body")
		return errors.ErrBadRequest
	}

	address.UserID = userID

	if err := database.DB.Create(&address).Error; err != nil {
		log.Print(err, "Failed to add address")
		return errors.ErrInternalServer
	}

	return c.Status(fiber.StatusCreated).JSON(address)
}

// GetUserAddresses - Fetch all addresses for the logged-in user
func GetUserAddresses(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)
	var addresses []models.Address

	if err := database.DB.Where("user_id = ?", userID).Find(&addresses).Error; err != nil {
		log.Print(err, "Failed to fetch addresses")
		return errors.ErrInternalServer
	}

	return c.JSON(fiber.Map{"addresses": addresses})
}

// DeleteAddress - Deletes an address (only if it belongs to the logged-in user)
func DeleteAddress(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)
	addressID := c.Params("id")

	var address models.Address
	if err := database.DB.First(&address, "id = ? AND user_id = ?", addressID, userID).Error; err != nil {
		log.Print(err, "Address not found or unauthorized access")
		return errors.ErrNotFound
	}

	if err := database.DB.Delete(&address).Error; err != nil {
		log.Print(err, "Failed to delete address")
		return errors.ErrInternalServer
	}

	return c.JSON(fiber.Map{"message": "Address deleted successfully"})
}
