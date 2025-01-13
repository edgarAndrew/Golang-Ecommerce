// controllers/user_controller.go
package controllers

import (
	"Golang-Rest-API/database"
	"Golang-Rest-API/errors"
	"Golang-Rest-API/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"log"
)

func UpdateUser(c *fiber.Ctx) error {
	var user models.User
	database.DB.First(&user, "id = ?", c.Locals("userID"))

	// Check if the user is an admin
	if !user.IsAdmin() {
		log.Print("User is not an admin")
		return errors.ErrForbidden
	}

	userID := c.Params("id")
	if userID == "" {
		log.Print("user ID is required")
		return errors.ErrBadRequest
	}

	type UpdateUserRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	request := new(UpdateUserRequest)

	if err := c.BodyParser(request); err != nil {
		log.Print(err, "Invalid request body")
		return errors.ErrBadRequest
	}

	obj := new(models.User)
	if err := database.DB.First(&obj, userID).Error; err != nil {
		log.Print(err, "Failed to find user")
		return errors.ErrNotFound
	}

	if request.Username != "" {
		obj.Username = request.Username
	}

	if request.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Print(err, "Failed to hash password")
			return errors.ErrInternalServer
		}
		obj.Password = string(hashedPassword)
	}

	if err := database.DB.Save(&obj).Error; err != nil {
		log.Print(err, "Failed to update user")
		return errors.ErrInternalServer
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
	})
}

func DeleteUser(c *fiber.Ctx) error {
	var user models.User
	database.DB.First(&user, "id = ?", c.Locals("userID"))

	// Check if the user is an admin
	if !user.IsAdmin() {
		log.Print("User is not an admin")
		return errors.ErrForbidden
	}

	userID := c.Params("id")
	if userID == "" {
		log.Print("user ID is required")
		return errors.ErrBadRequest
	}

	if err := database.DB.Delete(&models.User{}, userID).Error; err != nil {
		log.Print(err, "Failed to delete user")
		return errors.ErrInternalServer
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}

func GetUsers(c *fiber.Ctx) error {
	var user models.User
	database.DB.First(&user, "id = ?", c.Locals("userID"))

	// Check if the user is an admin
	if !user.IsAdmin() {
		log.Print("User is not an admin")
		return errors.ErrForbidden
	}

	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		log.Print(err, "Failed to fetch users")
		return errors.ErrInternalServer
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"users": users,
	})
}
