package models

import "gorm.io/gorm"

type Reservation struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	RoomID   uint   `gorm:"not null" json:"room_id"`
	PlayerID uint   `gorm:"not null" json:"player_id"`
	Date     string `gorm:"type:varchar(10);not null" json:"date"`
	Time     string `gorm:"type:varchar(5);not null" json:"time"`
	Player   Player `gorm:"foreignKey:PlayerID" json:"player"`
	Room     Room   `gorm:"foreignKey:RoomID" json:"room"`
}

func MigrateReservations(db *gorm.DB) error {
	// Migrate the Reservation table
	return db.AutoMigrate(&Reservation{})
}
