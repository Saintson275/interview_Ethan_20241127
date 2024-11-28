package main

import (
	"log"
	"os"

	"oxo/handlers"
	"oxo/models"
	"oxo/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Loading environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found, using system environment variables.")
	}

	// Init database connection
	db, err := storage.NewConnection(&storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Database Migration
	if err := models.Migrate(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Init Fiber
	app := fiber.New()

	// Registering Routes
	handlers.RegisterRoutes(app, db)

	// Start server
	log.Fatal(app.Listen(":8080"))
}
