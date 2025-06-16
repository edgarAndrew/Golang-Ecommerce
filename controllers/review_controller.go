package controllers

import (
	"Golang-Rest-API/database"
	"Golang-Rest-API/errors"
	"Golang-Rest-API/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

// AddReview - Allows users to review a product
func AddReview(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	review := new(models.Review)

	if err := c.BodyParser(review); err != nil {
		log.Print(err, "Invalid request body")
		return errors.ErrBadRequest
	}

	review.UserID = userID
	if err := database.DB.Create(&review).Error; err != nil {
		log.Print(err, "Failed to create review")
		return errors.ErrInternalServer
	}

	return c.Status(fiber.StatusCreated).JSON(review)
}

// GetReviewsForProduct - Fetch all reviews for a product
func GetReviewsForProduct(c *fiber.Ctx) error {
	productID := c.Params("productID")
	var reviews []models.Review

	// Query all reviews for the given ProductID
	if err := database.DB.Where("product_id = ?", productID).Find(&reviews).Error; err != nil {
		log.Print(err, "Failed to fetch reviews")
		return errors.ErrInternalServer
	}

	if len(reviews) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No reviews found for this product",
		})
	}

	return c.JSON(fiber.Map{"reviews": reviews})
}
