package models

import "time"

type Bed struct {
	ID        uint `gorm:"primaryKey"`
	RoomID    uint
	BedNumber string `gorm:"size:20;not null"`
	Status    string `gorm:"size:50;default:AVAILABLE"`
	CreatedAt time.Time
}
