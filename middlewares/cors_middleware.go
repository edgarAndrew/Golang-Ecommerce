package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// CORSConfig returns the CORS middleware with your custom configuration
func CORSConfig() fiber.Handler {
	return cors.New(cors.Config{
		// Allow all origins, you can set this to specific origins like "http://example.com"
		AllowOrigins: "http://localhost:3000",

		// Allow methods like GET, POST, PUT, DELETE, OPTIONS, etc.
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",

		// Allow specific headers such as Content-Type and Authorization
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",

		// Allow credentials such as cookies or HTTP authentication
		AllowCredentials: true,

		// MaxAge specifies how long the results of a preflight request can be cached by the browser
		MaxAge: 3600,
	})
}
