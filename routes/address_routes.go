package routes

import (
	"Golang-Rest-API/controllers"
	"Golang-Rest-API/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AddressRoutes(app *fiber.App) {
	app.Post("/api/addresses", middlewares.AuthMiddleware, controllers.AddAddress)
	app.Get("/api/addresses", middlewares.AuthMiddleware, controllers.GetUserAddresses)
	app.Delete("/api/addresses/:id", middlewares.AuthMiddleware, controllers.DeleteAddress)
}
