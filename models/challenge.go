package models

import (
	"time"

	"gorm.io/gorm"
)

// Challenge
type Challenge struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PlayerID  uint      `json:"player_id"`
	Amount    float64   `gorm:"default:20.01" json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}

// ChallengeResult
type ChallengeResult struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ChallengeID uint      `json:"challenge_id"`
	PlayerID    uint      `json:"player_id"`
	WinAmount   float64   `json:"win_amount"`
	Win         bool      `json:"win"`
	CreatedAt   time.Time `json:"created_at"`
}

func MigrateChallenges(db *gorm.DB) error {
	return db.AutoMigrate(&Challenge{}, &ChallengeResult{})
}
