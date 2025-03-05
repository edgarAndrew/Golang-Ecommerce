package routes

import (
	"Golang-Rest-API/controllers"
	"Golang-Rest-API/middlewares"

	"github.com/gofiber/fiber/v2"
)

func WishlistRoutes(app *fiber.App) {
	app.Post("/api/wishlist", middlewares.AuthMiddleware, controllers.AddToWishlist)
	app.Get("/api/wishlist", middlewares.AuthMiddleware, controllers.GetWishlist)
	app.Delete("/api/wishlist/:product_id", middlewares.AuthMiddleware, controllers.DeleteFromWishlist)
	app.Put("/api/wishlist/:product_id/increase", middlewares.AuthMiddleware, controllers.IncreaseWishlistCount)
	app.Put("/api/wishlist/:product_id/decrease", middlewares.AuthMiddleware, controllers.DecreaseWishlistCount)
}
