// middlewares/auth_middleware.go
package middlewares

import (
	"Golang-Rest-API/database"
	"Golang-Rest-API/errors"
	"Golang-Rest-API/models"
	"Golang-Rest-API/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware checks if the incoming request has a valid JWT token and attaches the userID to the context.
func AuthMiddleware(c *fiber.Ctx) error {
	// Get the token from the cookie
	cookie := c.Cookies("auth_token")
	if cookie == "" {
		log.Print("Missing authentication token")
		return errors.ErrUnauthorized
	}

	// Validate the token
	userID, err := utils.ValidateToken(cookie)
	if err != nil {
		log.Print(err, "Invalid token")
		return errors.ErrUnauthorized
	}

	// Attach userID to the context
	c.Locals("userID", userID)

	// Fetch the user from the database to ensure existence
	var user models.User
	if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
		log.Print(err, "Failed to find user")
		return errors.ErrNotFound
	}

	// Proceed to the next handler
	return c.Next()
}
