package utils

import (
	"hospital/models"
	"hospital/responses"
	"time"
)

// ======================
// Calculate Age
// ======================

func CalculateAge(dob time.Time) int {
	now := time.Now()

	age := now.Year() - dob.Year()

	// If birthday hasn't occurred yet this year
	if now.YearDay() < dob.YearDay() {
		age--
	}

	return age
}

// ======================
// Age Group
// ======================

func GetAgeGroup(age int) string {

	if age < 13 {
		return "Child"
	}

	if age < 18 {
		return "Teen"
	}

	if age < 60 {
		return "Adult"
	}

	return "Senior"
}

// ======================
// Mapper Function
// ======================

func MapPatientToResponse(p models.Patient) responses.PatientResponse {

	age := CalculateAge(p.DOB)

	fullName := p.FirstName
	if p.LastName != "" {
		fullName = p.FirstName + " " + p.LastName
	}

	return responses.PatientResponse{
		ID:        p.ID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		FullName:  fullName,
		Gender:    p.Gender,
		DOB:       p.DOB.Format("2006-01-02"),
		Age:       age,
		AgeGroup:  GetAgeGroup(age),
	}
}
