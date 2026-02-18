package models

import "time"

type Ward struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Type      string `gorm:"size:50;not null"` // ICU / GENERAL
	Floor     int
	CreatedAt time.Time
}
