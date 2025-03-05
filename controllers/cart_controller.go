package controllers

import (
	"Golang-Rest-API/database"
	"Golang-Rest-API/errors"
	"Golang-Rest-API/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

// AddToCart - Adds a product to the cart
func AddToCart(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	cart := new(models.Cart)

	if err := c.BodyParser(cart); err != nil {
		log.Print(err, "Invalid request body")
		return errors.ErrBadRequest
	}

	// Check if the product already exists in the user's cart
	var existingCart models.Cart
	if err := database.DB.Where("user_id = ? AND product_id = ?", userID, cart.ProductID).First(&existingCart).Error; err == nil {
		// If the product exists, increment the count
		existingCart.Count++
		if err := database.DB.Save(&existingCart).Error; err != nil {
			log.Print(err, "Failed to update cart")
			return errors.ErrInternalServer
		}
		return c.Status(fiber.StatusOK).JSON(existingCart)
	}

	// If the product doesn't exist, create a new cart entry
	cart.UserID = userID
	if err := database.DB.Create(&cart).Error; err != nil {
		log.Print(err, "Failed to add to cart")
		return errors.ErrInternalServer
	}

	return c.Status(fiber.StatusCreated).JSON(cart)
}

// GetUserCart - Fetches all items in the user's cart
func GetUserCart(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	var cartItems []models.Cart

	if err := database.DB.Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		log.Print(err, "Failed to fetch cart items")
		return errors.ErrInternalServer
	}

	return c.JSON(fiber.Map{"cart": cartItems})
}

// DeleteProductFromCart - Deletes a product from the user's cart
func DeleteProductFromCart(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	productID := c.Params("product_id")

	var cart models.Cart

	// Find the cart item by product ID and user ID
	if err := database.DB.First(&cart, "user_id = ? AND product_id = ?", userID, productID).Error; err != nil {
		log.Print(err, "Product not found in cart or unauthorized access")
		return errors.ErrNotFound
	}

	// Delete the cart item
	if err := database.DB.Delete(&cart).Error; err != nil {
		log.Print(err, "Failed to delete product from cart")
		return errors.ErrInternalServer
	}

	return c.JSON(fiber.Map{"message": "Product deleted from cart"})
}

// IncreaseCount - Increases the count of a product in the cart
func IncreaseCartCount(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	productID := c.Params("product_id")

	var cart models.Cart
	if err := database.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&cart).Error; err != nil {
		log.Print(err, "Product not found in cart")
		return errors.ErrNotFound
	}

	cart.Count++
	if err := database.DB.Save(&cart).Error; err != nil {
		log.Print(err, "Failed to increase count")
		return errors.ErrInternalServer
	}

	return c.Status(fiber.StatusOK).JSON(cart)
}

// DecreaseCount - Decreases the count of a product in the cart
func DecreaseCartCount(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	productID := c.Params("product_id")

	var cart models.Cart
	if err := database.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&cart).Error; err != nil {
		log.Print(err, "Product not found in cart")
		return errors.ErrNotFound
	}

	if cart.Count > 1 {
		cart.Count--
	} else {
		// Optionally, you can delete the product from the cart if the count reaches 0
		if err := database.DB.Delete(&cart).Error; err != nil {
			log.Print(err, "Failed to remove product from cart")
			return errors.ErrInternalServer
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Product removed from cart",
		})
	}

	if err := database.DB.Save(&cart).Error; err != nil {
		log.Print(err, "Failed to decrease count")
		return errors.ErrInternalServer
	}

	return c.Status(fiber.StatusOK).JSON(cart)
}
