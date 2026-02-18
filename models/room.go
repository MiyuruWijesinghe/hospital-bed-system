package models

import "time"

type Room struct {
	ID         uint `gorm:"primaryKey"`
	WardID     uint
	RoomNumber string `gorm:"size:20;not null"`
	RoomType   string `gorm:"size:50"` // PRIVATE / SHARED
	CreatedAt  time.Time
}
