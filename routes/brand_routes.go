package routes

import (
	"Golang-Rest-API/controllers"
	"Golang-Rest-API/middlewares"

	"github.com/gofiber/fiber/v2"
)

func BrandRoutes(app *fiber.App) {
	app.Post("/api/brands", middlewares.AuthMiddleware, controllers.CreateBrand)  // Admin-only
	app.Get("/api/brands/:id", controllers.GetBrand)
	app.Get("/api/brands", controllers.GetAllBrands)
	app.Delete("/api/brands/:id", middlewares.AuthMiddleware, controllers.DeleteBrand) // Admin-only
}
