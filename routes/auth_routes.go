// routes/auth_routes.go
package routes

import (
	"Golang-Rest-API/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	app.Post("/api/signup", controllers.Signup)
	app.Post("/api/login", controllers.Login)
}
