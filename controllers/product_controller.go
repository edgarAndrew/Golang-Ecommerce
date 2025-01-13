// controllers/product_controller.go
package controllers

import (
	"Golang-Rest-API/database"
	"Golang-Rest-API/errors"
	"Golang-Rest-API/models"
	"bytes"
	"context"
	"io"
	"log"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"

	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {
	var user models.User
	database.DB.First(&user, "id = ?", c.Locals("userID"))

	// Check if the user is an admin
	if !user.IsAdmin() {
		log.Print("User is not an admin")
		return errors.ErrForbidden
	}

	// Parse the product data from the request body
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		log.Print(err, "Invalid request body")
		return errors.ErrBadRequest
	}

	// Save the product to the database
	if err := database.DB.Create(&product).Error; err != nil {
		log.Print(err, "Failed to create product")
		return errors.ErrInternalServer
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product created successfully",
		"product": product,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	// Implementation here
	return nil
}

func DeleteProduct(c *fiber.Ctx) error {
	// Get user ID from context
	userID := c.Locals("userID")

	// Fetch the user to verify admin status
	var user models.User
	if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
		log.Print(err, "Failed to find user")
		return errors.ErrNotFound
	}

	// Check admin permissions
	if !user.IsAdmin() {
		log.Print("User is not an admin")
		return errors.ErrForbidden
	}

	// Get product ID from URL params
	productID := c.Params("id")
	if productID == "" {
		log.Print("Product ID is required")
		return errors.ErrBadRequest
	}

	// Find and delete the product
	if err := database.DB.Delete(&models.Product{}, productID).Error; err != nil {
		log.Print(err, "Failed to delete product")
		return errors.ErrInternalServer
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product

	// Fetch all products
	if err := database.DB.Find(&products).Error; err != nil {
		log.Print(err, "Failed to retrieve products")
		return errors.ErrInternalServer
	}

	var response []map[string]interface{}

	for _, prod := range products {
		var images []models.Image
		if err := database.DB.Where("product_id = ?", prod.ID).Find(&images).Error; err != nil {
			log.Printf("%s, Failed to fetch product image for product %v", err, prod.ID)
			// return errors.ErrInternalServer
		}
		imageURL := ""
		if len(images) > 0 {
			imageURL = images[0].Url
		}

		// Append product and image URL to the response
		response = append(response, map[string]interface{}{
			"product": prod,
			"image":   imageURL,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"products": response,
	})
}

func AddImageToProduct(c *fiber.Ctx) error {
	var user models.User
	database.DB.First(&user, "id = ?", c.Locals("userID"))

	// Check if the user is an admin
	if !user.IsAdmin() {
		log.Print("User is not an admin")
		return errors.ErrForbidden
	}

	// Get product ID from URL params
	productID := c.Params("id")
	if productID == "" {
		log.Print("Product ID is required")
		return errors.ErrBadRequest
	}

	// Check if the product exists
	var product models.Product
	if err := database.DB.First(&product, "id = ?", productID).Error; err != nil {
		log.Print(err, "Product not found")
		return errors.ErrNotFound
	}

	// Retrieve the file from the request
	file, err := c.FormFile("image")
	if err != nil {
		log.Print(err, "Failed to retrieve image from request")
		return errors.ErrBadRequest
	}

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		log.Print(err, "Failed to open image file")
		return errors.ErrInternalServer
	}
	defer src.Close()

	// Read the file content
	imageData, err := io.ReadAll(src)
	if err != nil {
		log.Print(err, "Failed to read image data")
		return errors.ErrInternalServer
	}

	// Create a new ProductImage instance
	obj := new(models.Image)

	var ctx = context.Background()
	resp, err := database.Cloudinary.Upload.Upload(ctx, bytes.NewReader(imageData), uploader.UploadParams{PublicID: file.Filename})

	if err != nil {
		log.Print(err, "Failed to upload image to cloudinary")
		return errors.ErrInternalServer
	}

	// Save the image to the database
	obj.ProductID = product.ID
	obj.Url = resp.URL
	obj.Name = product.Name + "_" + file.Filename

	// Save the image to the database
	if err := database.DB.Create(&obj).Error; err != nil {
		log.Print(err, "Failed to save image to database")
		return errors.ErrInternalServer
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Image added successfully",
	})
}

func GetProductImages(c *fiber.Ctx) error {
	// Get product ID from URL params
	productID := c.Params("id")
	if productID == "" {
		log.Print("Product ID is required")
		return errors.ErrBadRequest
	}

	// Check if the product exists
	var product models.Product
	if err := database.DB.First(&product, "id = ?", productID).Error; err != nil {
		log.Print(err, "Product not found")
		return errors.ErrNotFound
	}

	// Fetch all images associated with the product
	var productImages []models.Image
	if err := database.DB.Where("product_id = ?", productID).Find(&productImages).Error; err != nil {
		log.Print(err, "Failed to fetch product images")
		return errors.ErrInternalServer
	}

	images := make([]string, len(productImages))
	for i, img := range productImages {
		images[i] = img.Url
	}

	// Return the images as a response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"product_id": productID,
		"images":     images,
	})

}
