package models

import "time"

type Admission struct {
	ID           uint `gorm:"primaryKey"`
	PatientID    uint
	Patient      Patient `gorm:"foreignKey:PatientID"`
	BedID        uint
	Bed          Bed       `gorm:"foreignKey:BedID"`
	AdmittedAt   time.Time `gorm:"autoCreateTime"`
	DischargedAt *time.Time
	Status       string `gorm:"size:50;default:ACTIVE"`
	Reason       string `gorm:"size:255"`
}
