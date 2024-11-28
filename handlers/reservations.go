package handlers

import (
	"oxo/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterReservationRoutes(app fiber.Router, db *gorm.DB) {
	// Reservations API
	app.Get("/reservations", func(c *fiber.Ctx) error {
		roomID := c.Query("room_id")
		date := c.Query("date")
		limit := c.Query("limit")

		var reservations []models.Reservation
		query := db.Preload("Player").Preload("Room")

		// Filter by room ID if provided
		if roomID != "" {
			query = query.Where("room_id = ?", roomID)
		}

		// Filter by date if provided
		if date != "" {
			parsedDate, err := time.Parse("2006-01-02", date)
			if err != nil {
				return c.Status(400).JSON(fiber.Map{"error": "Invalid date format, use YYYY-MM-DD"})
			}
			query = query.Where("date = ?", parsedDate)
		}

		// Limit number of results if provided
		if limit != "" {
			limitInt, err := strconv.Atoi(limit)
			if err != nil {
				return c.Status(400).JSON(fiber.Map{"error": "Invalid limit format"})
			}
			query = query.Limit(limitInt)
		}

		if err := query.Find(&reservations).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch reservations"})
		}
		return c.JSON(reservations)
	})

	// Create a new reservation
	app.Post("/reservations", func(c *fiber.Ctx) error {
		reservation := new(models.Reservation)

		if err := c.BodyParser(reservation); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request, could not parse body", "details": err.Error()})
		}

		// Validation to ensure that room_id and player_id exist
		var room models.Room
		if err := db.First(&room, reservation.RoomID).Error; err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Room not found"})
		}

		var player models.Player
		if err := db.First(&player, reservation.PlayerID).Error; err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Player not found"})
		}

		// Validate date and time
		parsedDate, err := time.Parse("2006-01-02", reservation.Date)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid date format, must be YYYY-MM-DD"})
		}

		if parsedDate.Before(time.Now().AddDate(0, 0, 1)) {
			return c.Status(400).JSON(fiber.Map{"error": "Reservation date must be today or in the future"})
		}

		parsedTime, err := time.Parse("15:04", reservation.Time)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid time format, must be HH:mm"})
		}

		if parsedDate.Equal(time.Now().Truncate(24*time.Hour)) && parsedTime.Before(time.Now().Truncate(time.Hour)) {
			return c.Status(400).JSON(fiber.Map{"error": "Reservation time must be the current time or later today"})
		}

		// Create reservation
		if err := db.Create(reservation).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create reservation"})
		}

		return c.JSON(reservation.ID)
	})

}
