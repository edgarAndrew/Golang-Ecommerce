package routes

import (
	"Golang-Rest-API/controllers"
	"Golang-Rest-API/middlewares"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(app *fiber.App) {
	app.Post("/api/categories", middlewares.AuthMiddleware, controllers.CreateCategory)  // Admin-only
	app.Get("/api/categories", controllers.GetCategories)
	app.Delete("/api/categories/:id", middlewares.AuthMiddleware, controllers.DeleteCategory) // Admin-only
}
