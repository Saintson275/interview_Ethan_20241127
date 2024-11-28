package models

import (
	"time"

	"gorm.io/gorm"
)

// GameLog
type GameLog struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PlayerID  uint      `json:"player_id"`
	Action    string    `gorm:"type:varchar(100);not null" json:"action"`
	Timestamp time.Time `gorm:"default:current_timestamp" json:"timestamp"`
	Details   string    `gorm:"type:text" json:"details"`
}

// MigrateGameLogs
func MigrateGameLogs(db *gorm.DB) error {
	return db.AutoMigrate(&GameLog{})
}
