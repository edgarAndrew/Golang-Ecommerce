// controllers/auth_controller.go
package controllers

import (
	"Golang-Rest-API/database"
	"Golang-Rest-API/errors"
	"Golang-Rest-API/models"
	"Golang-Rest-API/utils"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *fiber.Ctx) error {

	// c.BodyParser(user) is a Fiber method that parses the JSON body of the incoming HTTP request and maps it to the user struct.
	// If the request body cannot be parsed (e.g., invalid JSON), it returns an error.

	// Parse request body
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		log.Print(err, "Invalid request body")
		return errors.ErrBadRequest
	}

	// Validate input
	if user.Username == "" || user.Password == "" || user.Email == "" {
		log.Print("Email, Username and Password are required")
		return errors.ErrBadRequest
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Print(err, "Failed to hash password")
		return errors.ErrInternalServer
	}
	user.Password = string(hashedPassword)

	// Save user to database
	if err := database.DB.Create(user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}

func Login(c *fiber.Ctx) error {
	// Parse request body
	loginData := new(models.User)
	if err := c.BodyParser(loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request",
		})
	}

	// Validate input
	if loginData.Email == "" || loginData.Password == "" {
		log.Print("Email and Password are required")
		return errors.ErrBadRequest
	}

	// Find user by email
	user := new(models.User)
	if err := database.DB.Where("email = ?", loginData.Email).First(user).Error; err != nil {
		log.Print(err, "Failed to find user")
		return errors.ErrNotFound
	}

	// Compare hashed passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		log.Print(err, "Wrong email/password")
		return errors.ErrUnauthorized
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		log.Print(err, "Failed to generate token")
		return errors.ErrInternalServer
	}

	// Set the token as an HTTP-only cookie
	c.Cookie(&fiber.Cookie{
		Name:     "auth_token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour), // Token valid for 1 day
		HTTPOnly: true,                           // Prevent JavaScript access
		Secure:   true,                           // Use HTTPS in production
		SameSite: "Strict",                       // Prevent CSRF attacks
	})

	// Use hx-redirect to redirect to /
	c.Set("HX-Redirect", "/")
	return c.JSON(fiber.Map{
		"message": "Login successful",
	})
}
