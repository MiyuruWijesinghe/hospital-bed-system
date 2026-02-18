package models

import "time"

type Admission struct {
	ID           uint `gorm:"primaryKey"`
	PatientID    uint
	BedID        uint
	AdmittedAt   time.Time `gorm:"autoCreateTime"`
	DischargedAt *time.Time
	Status       string `gorm:"size:50;default:ACTIVE"`
}
