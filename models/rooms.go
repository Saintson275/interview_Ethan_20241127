package models

import "gorm.io/gorm"

type Room struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	Description string `gorm:"type:varchar(100);not null" json:"description"`
	Status      string `gorm:"type:varchar(100);not null" json:"status"` // e.g："available", "occupied", "closed"
}

func MigrateRooms(db *gorm.DB) error {

	return db.AutoMigrate(&Room{})
}
