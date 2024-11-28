package handlers

import (
	"math/rand"
	"oxo/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterChallengeRoutes(app fiber.Router, db *gorm.DB) {
	// Challenges API

	// Players join the challenge
	app.Post("/challenges", func(c *fiber.Ctx) error {
		// Define a structure to receive player_id
		var request struct {
			PlayerID uint `json:"player_id"`
		}

		// Parsing the request body
		if err := c.BodyParser(&request); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		if request.PlayerID == 0 {
			return c.Status(400).JSON(fiber.Map{"error": "Player ID is required"})
		}

		// Check if you can join the challenge again (once per minute)
		lastChallenge := models.Challenge{}
		if err := db.Where("player_id = ?", request.PlayerID).Order("timestamp desc").First(&lastChallenge).Error; err != nil && err != gorm.ErrRecordNotFound {
			return c.Status(500).JSON(fiber.Map{"error": "Error checking last challenge"})
		}

		// If it is less than 1 minute since the last challenge, the challenge is forbidden
		if !lastChallenge.Timestamp.IsZero() && time.Now().Sub(lastChallenge.Timestamp) < time.Minute {
			return c.Status(400).JSON(fiber.Map{"error": "You can only participate once per minute"})
		}

		// Create a challenge record
		challenge := models.Challenge{
			PlayerID:  request.PlayerID,
			Amount:    20.01, // fixed amount
			Timestamp: time.Now(),
		}
		if err := db.Create(&challenge).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create challenge"})
		}

		// Winning is determined by a 1% chance
		winChance := rand.Float64() * 100
		win := winChance <= 1
		winAmount := 0.0
		if win {
			winAmount = 20.01
		}

		// Record challenge results
		result := models.ChallengeResult{
			ChallengeID: challenge.ID,
			PlayerID:    request.PlayerID,
			Win:         win,
			WinAmount:   winAmount,
			CreatedAt:   time.Now(),
		}
		if err := db.Create(&result).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create challenge result"})
		}

		// Return challenge results
		return c.JSON(fiber.Map{
			"challenge_id": challenge.ID,
			"status":       win,
			"win_amount":   winAmount,
		})
	})

	// View recent challenge results
	app.Get("/challenges/results", func(c *fiber.Ctx) error {
		var results []models.ChallengeResult
		if err := db.Order("created_at desc").Limit(10).Find(&results).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch challenge results"})
		}
		return c.JSON(results)
	})
}
