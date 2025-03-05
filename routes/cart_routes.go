package routes

import (
	"Golang-Rest-API/controllers"
	"Golang-Rest-API/middlewares"

	"github.com/gofiber/fiber/v2"
)

func CartRoutes(app *fiber.App) {
	app.Post("/api/cart", middlewares.AuthMiddleware, controllers.AddToCart)
	app.Get("/api/cart", middlewares.AuthMiddleware, controllers.GetUserCart)
	app.Delete("/api/cart/:product_id", middlewares.AuthMiddleware, controllers.DeleteProductFromCart)
	app.Put("/api/cart/:product_id/increase", middlewares.AuthMiddleware, controllers.IncreaseCartCount)
	app.Put("/api/cart/:product_id/decrease", middlewares.AuthMiddleware, controllers.DecreaseCartCount)
}
