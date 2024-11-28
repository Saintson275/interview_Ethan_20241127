package models

import (
	"encoding/json"
	"log"
	"time"

	"gorm.io/gorm"
)

// Payment - 支付模型
type Payment struct {
	ID            uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	PlayerID      uint            `json:"player_id"`
	Method        string          `json:"method"`
	Amount        float64         `json:"amount"`
	Details       json.RawMessage `json:"details"`
	TransactionID string          `json:"transaction_id"`
	Status        string          `json:"status"`
	CreatedAt     time.Time       `json:"created_at"`
}

func MigratePayments(db *gorm.DB) error {
	err := db.AutoMigrate(&Payment{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate Payment model: %v", err)
	}
	return err
}
