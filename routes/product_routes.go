package routes

import (
	"Golang-Rest-API/controllers"
	"Golang-Rest-API/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App) {
	app.Post("/api/products", middlewares.AuthMiddleware, controllers.CreateProduct)
	app.Put("/api/products/:id", middlewares.AuthMiddleware, controllers.UpdateProduct)
	app.Delete("/api/products/:id", middlewares.AuthMiddleware, controllers.DeleteProduct)
	app.Get("/api/products", controllers.GetProducts)                    // Get all products with limited info
	app.Get("/api/products/:id", controllers.GetProductDetails)          // Get detailed product information

	app.Post("/api/products/add-image/:id", middlewares.AuthMiddleware, controllers.AddImageToProduct)
	app.Get("/api/products/get-image/:id", controllers.GetProductImages)
}
