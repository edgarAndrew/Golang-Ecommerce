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

// CreateProduct creates a new product
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

// UpdateProduct updates an existing product
func UpdateProduct(c *fiber.Ctx) error {
	// Get user ID from context
	userID := c.Locals("userID")
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

	// Find the product
	var product models.Product
	if err := database.DB.First(&product, "id = ?", productID).Error; err != nil {
		log.Print(err, "Product not found")
		return errors.ErrNotFound
	}

	// Parse the new data
	if err := c.BodyParser(&product); err != nil {
		log.Print(err, "Failed to parse product data")
		return errors.ErrBadRequest
	}

	// Save the updated product
	if err := database.DB.Save(&product).Error; err != nil {
		log.Print(err, "Failed to update product")
		return errors.ErrInternalServer
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product updated successfully",
		"product": product,
	})
}

// DeleteProduct deletes a product
func DeleteProduct(c *fiber.Ctx) error {
	// Get user ID from context
	userID := c.Locals("userID")
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

// GetProducts retrieves all products with basic information and one image
func GetProducts(c *fiber.Ctx) error {
    var products []models.Product

    // Preload Images and Reviews when fetching products
    if err := database.DB.Preload("Images").Preload("Reviews").Find(&products).Error; err != nil {
        log.Print(err, "Failed to retrieve products")
        return errors.ErrInternalServer
    }

    var response []map[string]interface{}

    // Loop through products and retrieve associated images
    for _, prod := range products {
        // Ensure only one image is included (using the first image, if available)
        imageURL := ""
        if len(prod.Images) > 0 {
            imageURL = prod.Images[0].Url
        }

        // Append product, image URL, and reviews to the response
        response = append(response, map[string]interface{}{
            "ID":                 prod.ID,
            "BrandID":            prod.BrandID,
            "SubCategoryID":      prod.SubCategoryID,
            "ProductName":        prod.ProductName,
            "ProductDescription": prod.ProductDescription,
            "Price":              prod.Price,
            "Stock":              prod.Stock,
            "SalePercentage":     prod.SalePercentage,
            "Image":              imageURL,  // Only one image
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "products": response,
    })
}


// GetProductDetails retrieves detailed information for a specific product
func GetProductDetails(c *fiber.Ctx) error {
	productID := c.Params("id")
    if productID == "" {
        log.Print("Product ID is required")
        return errors.ErrBadRequest
    }

    // Find the product by ID
    var product models.Product
    if err := database.DB.Preload("Images").Preload("Reviews").First(&product, "id = ?", productID).Error; err != nil {
        log.Print(err, "Product not found")
        return errors.ErrNotFound
    }

    // Return the product with all details
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "product": product,
    })
}


// AddImageToProduct adds an image to a product
func AddImageToProduct(c *fiber.Ctx) error {
	var user models.User
	database.DB.First(&user, "id = ?", c.Locals("userID"))

	// Check if the user is an admin
	if !user.Admin {
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
	obj.Name = product.ProductName + "_" + file.Filename

	// Save the image to the database
	if err := database.DB.Create(&obj).Error; err != nil {
		log.Print(err, "Failed to save image to database")
		return errors.ErrInternalServer
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Image added successfully",
	})
}

// GetProductImages retrieves images for a specific product
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
