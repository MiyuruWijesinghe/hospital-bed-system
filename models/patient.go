package models

import "time"

type Patient struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"size:100;not null"`
	LastName  string `gorm:"size:100"`
	Gender    string `gorm:"size:20"`
	DOB       time.Time
	CreatedAt time.Time
}
