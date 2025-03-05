package routes

import (
	"Golang-Rest-API/controllers"
	"Golang-Rest-API/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ReviewRoutes(app *fiber.App) {
	app.Post("/api/reviews", middlewares.AuthMiddleware, controllers.AddReview)
	app.Get("/api/products/:product_id/reviews", controllers.GetReviewsForProduct)
}
