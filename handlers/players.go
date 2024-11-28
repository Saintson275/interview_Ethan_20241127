package handlers

import (
	"oxo/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterPlayerRoutes(app fiber.Router, db *gorm.DB) {
	// Players API
	app.Get("/players", func(c *fiber.Ctx) error {
		var players []models.Player
		if err := db.Find(&players).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch players"})
		}
		return c.JSON(players)
	})

	app.Post("/players", func(c *fiber.Ctx) error {
		player := new(models.Player)
		if err := c.BodyParser(player); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}
		if err := db.Create(player).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create player"})
		}
		return c.JSON(player)
	})

	app.Get("/players/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var player models.Player
		if err := db.First(&player, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Player not found"})
		}
		return c.JSON(player)
	})

	app.Put("/players/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var player models.Player
		if err := db.First(&player, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Player not found"})
		}
		if err := c.BodyParser(&player); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}
		db.Save(&player)
		return c.JSON(player)
	})

	app.Delete("/players/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		if err := db.Delete(&models.Player{}, id).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to delete player"})
		}
		return c.SendStatus(204)
	})
}
