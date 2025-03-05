package controllers

import (
	"Golang-Rest-API/database"
	"Golang-Rest-API/errors"
	"Golang-Rest-API/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

// AddToWishlist - Adds a product to the user's wishlist
func AddToWishlist(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	wishlist := new(models.Wishlist)

	if err := c.BodyParser(wishlist); err != nil {
		log.Print(err, "Invalid request body")
		return errors.ErrBadRequest
	}

	// Check if the product already exists in the wishlist
	var existingWishlist models.Wishlist
	if err := database.DB.Where("user_id = ? AND product_id = ?", userID, wishlist.ProductID).First(&existingWishlist).Error; err == nil {
		// If product exists, increment the count
		existingWishlist.Count++
		if err := database.DB.Save(&existingWishlist).Error; err != nil {
			log.Print(err, "Failed to update wishlist")
			return errors.ErrInternalServer
		}
		return c.Status(fiber.StatusOK).JSON(existingWishlist)
	}

	// If the product does not exist in the wishlist, create a new entry
	wishlist.UserID = userID
	if err := database.DB.Create(&wishlist).Error; err != nil {
		log.Print(err, "Failed to add to wishlist")
		return errors.ErrInternalServer
	}

	return c.Status(fiber.StatusCreated).JSON(wishlist)
}

// GetWishlist - Retrieves all products in the user's wishlist
func GetWishlist(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	var wishlist []models.Wishlist

	if err := database.DB.Where("user_id = ?", userID).Find(&wishlist).Error; err != nil {
		log.Print(err, "Failed to fetch wishlist")
		return errors.ErrInternalServer
	}

	if len(wishlist) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Wishlist is empty",
		})
	}

	return c.JSON(fiber.Map{"wishlist": wishlist})
}

// DeleteFromWishlist - Deletes a product from the user's wishlist
func DeleteFromWishlist(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	productID := c.Params("productID")

	var wishlist models.Wishlist
	if err := database.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&wishlist).Error; err != nil {
		log.Print(err, "Product not found in wishlist")
		return errors.ErrNotFound
	}

	if err := database.DB.Delete(&wishlist).Error; err != nil {
		log.Print(err, "Failed to remove product from wishlist")
		return errors.ErrInternalServer
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product removed from wishlist",
	})
}

// IncreaseCount - Increases the count of a product in the wishlist
func IncreaseWishlistCount(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	productID := c.Params("productID")

	var wishlist models.Wishlist
	if err := database.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&wishlist).Error; err != nil {
		log.Print(err, "Product not found in wishlist")
		return errors.ErrNotFound
	}

	wishlist.Count++
	if err := database.DB.Save(&wishlist).Error; err != nil {
		log.Print(err, "Failed to increase count")
		return errors.ErrInternalServer
	}

	return c.Status(fiber.StatusOK).JSON(wishlist)
}

// DecreaseCount - Decreases the count of a product in the wishlist
func DecreaseWishlistCount(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	productID := c.Params("productID")

	var wishlist models.Wishlist
	if err := database.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&wishlist).Error; err != nil {
		log.Print(err, "Product not found in wishlist")
		return errors.ErrNotFound
	}

	if wishlist.Count > 1 {
		wishlist.Count--
	} else {
		// Optionally, you can delete the product from the wishlist if the count reaches 0
		if err := database.DB.Delete(&wishlist).Error; err != nil {
			log.Print(err, "Failed to remove product from wishlist")
			return errors.ErrInternalServer
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Product removed from wishlist",
		})
	}

	if err := database.DB.Save(&wishlist).Error; err != nil {
		log.Print(err, "Failed to decrease count")
		return errors.ErrInternalServer
	}

	return c.Status(fiber.StatusOK).JSON(wishlist)
}
