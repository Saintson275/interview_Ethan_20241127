package handlers

import (
	"oxo/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterLogRoutes(app fiber.Router, db *gorm.DB) {
	// Logs API
	app.Get("/logs", func(c *fiber.Ctx) error {
		playerID := c.Query("player_id")
		action := c.Query("action")
		startTime := c.Query("start_time")
		endTime := c.Query("end_time")
		limit := c.Query("limit")

		var logs []models.GameLog
		query := db.Model(&models.GameLog{})

		// Filter by player ID if provided
		if playerID != "" {
			query = query.Where("player_id = ?", playerID)
		}

		// Filter by action if provided
		if action != "" {
			query = query.Where("action = ?", action)
		}

		// Filter by start time and end time if provided
		if startTime != "" && endTime != "" {
			parsedStartTime, err := time.Parse("2006-01-02", startTime)
			if err != nil {
				return c.Status(400).JSON(fiber.Map{"error": "Invalid start_time format"})
			}
			parsedEndTime, err := time.Parse("2006-01-02", endTime)
			if err != nil {
				return c.Status(400).JSON(fiber.Map{"error": "Invalid end_time format"})
			}
			query = query.Where("timestamp BETWEEN ? AND ?", parsedStartTime, parsedEndTime)
		}

		// Limit number of results if provided
		if limit != "" {
			limitInt, err := strconv.Atoi(limit)
			if err != nil {
				return c.Status(400).JSON(fiber.Map{"error": "Invalid limit format"})
			}
			query = query.Limit(limitInt)
		}

		if err := query.Find(&logs).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch logs"})
		}

		return c.JSON(logs)
	})

	app.Post("/logs", func(c *fiber.Ctx) error {
		log := new(models.GameLog)
		if err := c.BodyParser(log); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}
		if err := db.Create(log).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create log"})
		}
		return c.JSON(log)
	})
}
