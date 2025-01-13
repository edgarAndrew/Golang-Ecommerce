package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envVars := map[string]string{
		"PORT":                  os.Getenv("PORT"),
		"JWT_SECRET":            os.Getenv("JWT_SECRET"),
		"JWT_EXPIRATION_HOURS":  os.Getenv("JWT_EXPIRATION_HOURS"),
		"CLOUDINARY_NAME":       os.Getenv("CLOUDINARY_NAME"),
		"CLOUDINARY_API_KEY":    os.Getenv("CLOUDINARY_API_KEY"),
		"CLOUDINARY_API_SECRET": os.Getenv("CLOUDINARY_API_SECRET"),
		"ENABLE_CORS":           os.Getenv("ENABLE_CORS"),
		"USE_SQLITE_DB":         os.Getenv("USE_SQLITE_DB"),
		"SHOW_SQL":              os.Getenv("SHOW_SQL"),
	}

	missingVars := []string{}
	for key, value := range envVars {
		if value == "" {
			missingVars = append(missingVars, key)
		}
	}

	if len(missingVars) > 0 {
		log.Fatalf("Missing environment variables: %v", missingVars)
	}

	if envVars["SHOW_SQL"] != "true" && envVars["SHOW_SQL"] != "false" {
		log.Fatal("SHOW_SQL must be either 'true' or 'false'.")
	}

	if envVars["USE_SQLITE_DB"] != "true" && envVars["USE_SQLITE_DB"] != "false" {
		log.Fatal("USE_SQLITE_DB must be either 'true' or 'false'.")
	}

	if envVars["ENABLE_CORS"] != "true" && envVars["ENABLE_CORS"] != "false" {
		log.Fatal("ENABLE_CORS must be either 'true' or 'false'.")
	}

	if envVars["USE_SQLITE_DB"] != "true" {
		MySqlURI := os.Getenv("MYSQL_URI")
		if MySqlURI == "" {
			log.Fatal("MYSQL_URI is not set in env.")
		}
	}
}
