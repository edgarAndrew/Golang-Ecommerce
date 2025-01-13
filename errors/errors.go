package errors

import "github.com/gofiber/fiber/v2"

// Define reusable Fiber errors
var (
	ErrBadRequest     = fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	ErrUnauthorized   = fiber.NewError(fiber.StatusUnauthorized, "Unauthorized Access")
	ErrForbidden      = fiber.NewError(fiber.StatusForbidden, "Forbidden")
	ErrNotFound       = fiber.NewError(fiber.StatusNotFound, "Resource Not Found")
	ErrInternalServer = fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	ErrConflict       = fiber.NewError(fiber.StatusConflict, "Conflict")
	ErrUnprocessable  = fiber.NewError(fiber.StatusUnprocessableEntity, "Unprocessable Entity")
)
