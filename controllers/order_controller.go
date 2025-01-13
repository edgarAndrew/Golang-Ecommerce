// controllers/order_controller.go
package controllers

import (
	"Golang-Rest-API/database"
	"Golang-Rest-API/errors"
	"Golang-Rest-API/models"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateOrder(c *fiber.Ctx) error {
	var user models.User
	database.DB.First(&user, "id = ?", c.Locals("userID"))

	// Parse the order data from the request body
	type OrderCreateRequest struct {
		ProductID uint `json:"productId"`
		Quantity  int  `json:"quantity"`
	}
	order := new(OrderCreateRequest)

	if err := c.BodyParser(order); err != nil {
		log.Print(err, "Invalid request body")
		return errors.ErrBadRequest
	}

	var product models.Product
	database.DB.First(&product, "id = ?", order.ProductID)

	if product.ID == 0 {
		log.Print("Product not found")
		return errors.ErrNotFound
	}

	obj := new(models.Order)
	obj.ProductID = order.ProductID
	obj.Quantity = order.Quantity
	obj.UserID = user.ID
	obj.Status = "Pending"
	obj.User = user
	obj.Product = product

	// Save the order to the database
	if err := database.DB.Create(&obj).Error; err != nil {
		log.Print(err, "Failed to create order")
		return errors.ErrInternalServer
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Order created successfully",
		"order":   order,
	})
}

func ChangeOrderStatus(c *fiber.Ctx) error {
	var user models.User
	database.DB.First(&user, "id = ?", c.Locals("userID"))

	// Check admin permissions
	if !user.IsAdmin() {
		log.Print("User is not an admin")
		return errors.ErrForbidden
	}

	// Get order ID from URL params
	orderID := c.Params("id")
	if orderID == "" {
		log.Print("Order ID is required")
		return errors.ErrBadRequest
	}

	// Parse the order data from the request body
	type OrderStatusRequest struct {
		Status string `json:"status"`
	}

	order := new(OrderStatusRequest)
	if err := c.BodyParser(order); err != nil {
		log.Print(err, "Invalid request body")
		return errors.ErrBadRequest
	}

	// Save the order to the database
	obj := new(models.Order)
	database.DB.First(&obj, "id = ?", orderID)

	if obj.ID == 0 {
		log.Print("Order not found")
		return errors.ErrNotFound
	}

	obj.Status = order.Status

	if err := database.DB.Save(&obj).Error; err != nil {
		log.Print(err, "Failed to update order due to invalid status")
		return errors.ErrBadRequest
	}

	// if err := database.DB.Model(&models.Order{}).Where("id = ?", orderID).Update("status", order.Status).Error; err != nil {
	// 	log.Print(err, "Failed to update order due to invalid status or invalid order ID")
	// 	return errors.ErrBadRequest
	// }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Order status updated successfully",
	})
}

func GetAllOrders(c *fiber.Ctx) error {

	type OrderResponse struct {
		ID        uint      `json:"id"`
		UserID    uint      `json:"user_id"`
		ProductID uint      `json:"product_id"`
		Quantity  int       `json:"quantity"`
		Status    string    `json:"status"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	// Fetch the current user
	var user models.User
	database.DB.First(&user, "id = ?", c.Locals("userID"))

	// Check admin permissions
	if !user.IsAdmin() {
		log.Print("User is not an admin")
		return errors.ErrForbidden
	}

	// Fetch all orders
	var orders []models.Order
	if err := database.DB.Find(&orders).Error; err != nil {
		log.Print(err, "Failed to fetch orders")
		return errors.ErrInternalServer
	}

	// Map orders to the response struct
	response := make([]OrderResponse, len(orders))
	for i, order := range orders {
		response[i] = OrderResponse{
			ID:        order.ID,
			UserID:    order.UserID,
			ProductID: order.ProductID,
			Quantity:  order.Quantity,
			Status:    order.Status,
			CreatedAt: order.CreatedAt,
			UpdatedAt: order.UpdatedAt,
		}
	}

	// Return the response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"orders": response,
	})
}

func GetMyOrders(c *fiber.Ctx) error {
	var user models.User
	database.DB.First(&user, "id = ?", c.Locals("userID"))

	type OrderResponse struct {
		ID        uint      `json:"id"`
		UserID    uint      `json:"user_id"`
		ProductID uint      `json:"product_id"`
		Quantity  int       `json:"quantity"`
		Status    string    `json:"status"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	// Fetch all orders for the current user
	var orders []models.Order
	if err := database.DB.Where("user_id = ?", user.ID).Find(&orders).Error; err != nil {
		log.Print(err, "Failed to fetch orders")
		return errors.ErrInternalServer
	}

	response := make([]OrderResponse, len(orders))
	for i, order := range orders {
		response[i] = OrderResponse{
			ID:        order.ID,
			UserID:    order.UserID,
			ProductID: order.ProductID,
			Quantity:  order.Quantity,
			Status:    order.Status,
			CreatedAt: order.CreatedAt,
			UpdatedAt: order.UpdatedAt,
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"orders": response,
	})
}
