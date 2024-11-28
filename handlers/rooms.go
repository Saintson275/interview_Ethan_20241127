package handlers

import (
	"oxo/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoomRoutes(app fiber.Router, db *gorm.DB) {
	// Rooms API
	app.Get("/rooms", func(c *fiber.Ctx) error {
		var rooms []models.Room
		if err := db.Find(&rooms).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch rooms"})
		}
		return c.JSON(rooms)
	})

	app.Post("/rooms", func(c *fiber.Ctx) error {
		room := new(models.Room)
		if err := c.BodyParser(room); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}
		if err := db.Create(room).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create room"})
		}
		return c.JSON(room)
	})

	app.Get("/rooms/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var room models.Room
		if err := db.First(&room, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Room not found"})
		}
		return c.JSON(room)
	})

	app.Put("/rooms/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var room models.Room
		if err := db.First(&room, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Room not found"})
		}
		if err := c.BodyParser(&room); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}
		db.Save(&room)
		return c.JSON(room)
	})

	app.Delete("/rooms/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		if err := db.Delete(&models.Room{}, id).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to delete room"})
		}
		return c.SendStatus(204)
	})
}
