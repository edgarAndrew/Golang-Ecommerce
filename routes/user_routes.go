// routes/user_routes.go
package routes

import (
	"Golang-Rest-API/controllers"
	"Golang-Rest-API/middlewares"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Put("/api/users/:id", middlewares.AuthMiddleware, controllers.UpdateUser)
	app.Delete("/api/users/:id", middlewares.AuthMiddleware, controllers.DeleteUser)
	app.Get("/api/users", middlewares.AuthMiddleware, controllers.GetUsers)
}
