// middlewares/error_middleware.go
package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	// Check if the error is a *fiber.Error type
	if e, ok := err.(*fiber.Error); ok {
		// Return the specific HTTP error status and message
		return c.Status(e.Code).JSON(fiber.Map{
			"message": e.Message,
		})
	}

	// Fallback for unexpected errors
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "Internal Server Error",
		"details": err.Error(), // Optional: include details for debugging
	})
}

// NotFoundHandler handles 404 Not Found errors.
func NotFoundHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Resource not found",
	})
}
