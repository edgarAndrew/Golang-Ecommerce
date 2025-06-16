package main

import (
	"log"
	"os"

	"Golang-Rest-API/database"
	"Golang-Rest-API/middlewares"
	"Golang-Rest-API/routes"
	"Golang-Rest-API/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load environment variables
	utils.LoadEnv()

	// Create a new Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
	})

	// Serve static files (CSS, JS, Images, etc.)
	app.Static("/static", "./static")

	// Initialize the database
	database.ConnectDB()

	// Connect to Cloudinary
	database.ConnectCloudinary()

	// Middleware setup
	if os.Getenv("ENABLE_CORS") == "true" {
		app.Use(middlewares.CORSConfig())
	}

	// Set up routes
	routes.AddressRoutes(app)
	routes.AuthRoutes(app)
	routes.BrandRoutes(app)
	routes.CartRoutes(app)
	routes.CategoryRoutes(app)
	routes.OrderRoutes(app)
	routes.ProductRoutes(app)
	routes.ReviewRoutes(app)
	routes.SubCategoryRoutes(app)
	routes.UserRoutes(app)
	routes.WishlistRoutes(app)

	app.Use(middlewares.NotFoundHandler)

	// Start the server
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
	log.Printf("Server is running on port %s", os.Getenv("PORT"))
}
