package models

import "gorm.io/gorm"

type Level struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(100);not null" json:"name"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Player{}, &Level{})
}
