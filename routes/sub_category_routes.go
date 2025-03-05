package routes

import (
	"Golang-Rest-API/controllers"
	"Golang-Rest-API/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SubCategoryRoutes(app *fiber.App) {
	app.Post("/api/subcategories", middlewares.AuthMiddleware, controllers.CreateSubCategory)       // Admin-only
	app.Delete("/api/subcategories/:id", middlewares.AuthMiddleware, controllers.DeleteSubCategory) // Admin-only
	app.Get("/api/subcategories/:category_id", middlewares.AuthMiddleware, controllers.GetSubCategoriesByCategoryID)
}
