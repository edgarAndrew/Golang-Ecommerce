// utils/jwt.go
package utils

import (
	"Golang-Rest-API/errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

// GenerateToken creates a new JWT token for the user.
func GenerateToken(userID uint) (string, error) {
	expirationStr := os.Getenv("JWT_EXPIRATION_HOURS")
	if expirationStr == "" {
		expirationStr = "24" // Default
	}

	// Parse the expiration duration
	expirationHours, err := time.ParseDuration(expirationStr + "h")
	if err != nil {
		log.Print("invalid JWT_EXPIRATION_HOURS format")
		return "", errors.ErrBadRequest
	}

	// Create the claims, specifying the user ID and an expiration time.
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(expirationHours).Unix(), // Token expiration
	}

	// Create a new token with the claims and sign it with the secret key.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate the signed token as a string.
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken checks the validity of the given JWT token.
func ValidateToken(tokenString string) (int, error) {
	// Split the token into parts to extract the token string part.
	tokenString = strings.TrimSpace(tokenString)
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Parse the token.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token is signed with our secret key.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return -1, errors.ErrUnauthorized
		}
		return secretKey, nil
	})

	if err != nil {
		return -1, err
	}

	// Check if the token is valid and not expired.
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return -1, errors.ErrUnauthorized
	}

	// Extract the user ID from the claims.
	userID, ok := claims["userID"].(float64)
	if !ok {
		return -1, errors.ErrUnprocessable
	}

	// Return the user ID.
	return int(userID), nil
}
