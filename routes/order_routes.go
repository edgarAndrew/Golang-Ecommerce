// routes/order_routes.go
package routes

import (
	"Golang-Rest-API/controllers"
	"Golang-Rest-API/middlewares"

	"github.com/gofiber/fiber/v2"
)

func OrderRoutes(app *fiber.App) {
	app.Post("/api/orders", middlewares.AuthMiddleware, controllers.CreateOrder)
	app.Get("/api/orders", middlewares.AuthMiddleware, controllers.GetAllOrders)
	app.Get("/api/myorders", middlewares.AuthMiddleware, controllers.GetMyOrders)
	app.Patch("/api/orders/:id", middlewares.AuthMiddleware, controllers.ChangeOrderStatus)
}
