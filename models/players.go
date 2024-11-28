package models

import "gorm.io/gorm"

type Player struct {
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"type:varchar(100);not null" json:"name"`
	Level string `gorm:"not null" json:"level"`
}

func MigratePlayers(db *gorm.DB) error {
	err := db.AutoMigrate(&Player{})
	return err
}
