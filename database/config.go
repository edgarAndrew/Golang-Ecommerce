package database

import (
	"log"

	"Golang-Rest-API/models"

	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB         *gorm.DB // database connection pointer
	Err        error
	Cloudinary *cloudinary.Cloudinary // cloudinary connection pointer
)

func ConnectDB() {
	useLocalDB := os.Getenv("USE_SQLITE_DB")

	log.Print("Connecting to database ...")

	if useLocalDB == "true" {
		ConnectSQLiteDB()
	} else {
		ConnectMySQLDB()
	}
}

func ConnectCloudinary() {
	log.Print("Connecting to cloudinary ...")
	Cloudinary, Err = cloudinary.NewFromParams(os.Getenv("CLOUDINARY_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET"))
	if Err != nil {
		log.Fatalf("failed to connect to cloudinary: %v", Err)
	}
	log.Print("Connected to cloudinary.")
}

func ConnectMySQLDB() {
	dsn := os.Getenv("MYSQL_URI")

	showSQL := os.Getenv("SHOW_SQL") == "true"
	logLevel := logger.Silent
	if showSQL {
		logLevel = logger.Info
	}

	// Set GORM logger with the determined log level
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: true,
		},
	)

	DB, Err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger, // Use the configured logger
	})
	if Err != nil {
		log.Fatalf("failed to connect to database: %v", Err)
	}

	// Auto migrate models
	err := DB.AutoMigrate(
		&models.Address{},
		&models.Brand{},
		&models.Cart{},
		&models.Category{},
		&models.Image{},
		&models.OrderItem{},
		&models.Order{},
		&models.Product{},
		&models.Review{},
		&models.SubCategory{},
		&models.User{},
		&models.Wishlist{},
	)
	if err != nil {
		log.Fatalf("failed to Auto migrate models to database: %v", err)
	}
	log.Println("Auto migrated models.%v", err)

	log.Println("Connected to MySQL database.")
}

func ConnectSQLiteDB() {
	DB, Err = gorm.Open(sqlite.Open("sqlite.db?_foreign_keys=on"), &gorm.Config{})
	if Err != nil {
		// Fatal errors call the os.Exit(1) function and halts the main function
		log.Fatalf("failed to connect to database: %v", Err)
	}

	// Auto migrate models
	err := DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.Image{})
	if err != nil {
		log.Fatalf("failed to Auto migrate models to database: %v", err)
	}

	log.Println("Connected to SQLite database.")
}
