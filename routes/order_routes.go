package routes

import (
	"Golang-Rest-API/controllers"
	"Golang-Rest-API/middlewares"

	"github.com/gofiber/fiber/v2"
)

func OrderRoutes(app *fiber.App) {
	app.Post("/api/orders", middlewares.AuthMiddleware, controllers.CreateOrder)
	app.Get("/api/orders", middlewares.AuthMiddleware, controllers.GetMyOrders)
	app.Get("/api/orders/all", middlewares.AuthMiddleware, controllers.GetAllOrders) // For admins
	app.Put("/api/orders/:id", middlewares.AuthMiddleware, controllers.ChangeOrderStatus)  // New update API
}
