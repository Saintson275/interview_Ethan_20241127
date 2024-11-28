package handlers

import (
	"oxo/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterLevelRoutes(app fiber.Router, db *gorm.DB) {
	app.Get("/levels", func(c *fiber.Ctx) error {
		var levels []models.Level
		if err := db.Find(&levels).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch levels"})
		}
		return c.JSON(levels)
	})

	app.Post("/levels", func(c *fiber.Ctx) error {
		level := new(models.Level)
		if err := c.BodyParser(level); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}
		if err := db.Create(level).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create level"})
		}
		return c.JSON(level)
	})
}
