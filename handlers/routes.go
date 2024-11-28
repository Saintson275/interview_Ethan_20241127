package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// RegisterRoutes integrates all routes
func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")

	RegisterPlayerRoutes(api, db)

	RegisterLevelRoutes(api, db)

	RegisterReservationRoutes(api, db)

	RegisterRoomRoutes(api, db)

	RegisterChallengeRoutes(api, db)

	RegisterLogRoutes(api, db)

	RegisterPaymentRoutes(api, db)
}
