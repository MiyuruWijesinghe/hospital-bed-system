package controllers

import (
	"hospital/config"
	"hospital/models"
	"hospital/requests"
	"hospital/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ======================
// Calculate Age
// ======================

func calculateAge(dob time.Time) int {
	now := time.Now()

	age := now.Year() - dob.Year()

	// If birthday hasn't occurred yet this year
	if now.YearDay() < dob.YearDay() {
		age--
	}

	return age
}

// ======================
// Mapper Function
// ======================

func mapPatientToResponse(p models.Patient) responses.PatientResponse {
	return responses.PatientResponse{
		ID:        p.ID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Gender:    p.Gender,
		DOB:       p.DOB.Format("2006-01-02"),
		Age:       calculateAge(p.DOB),
	}
}

func CreatePatient(c *gin.Context) {
	var req requests.CreatePatientRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	patient := models.Patient{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Gender:    req.Gender,
		DOB:       dob,
	}

	config.DB.Create(&patient)

	response := mapPatientToResponse(patient)

	c.JSON(http.StatusCreated, response)
}

func GetPatients(c *gin.Context) {
	var patients []models.Patient
	config.DB.Find(&patients)

	var response []responses.PatientResponse

	for _, p := range patients {
		response = append(response, mapPatientToResponse(p))
	}

	c.JSON(http.StatusOK, response)
}

func GetPatientByID(c *gin.Context) {
	patientID := c.Param("patient_id")

	var patient models.Patient

	if err := config.DB.First(&patient, patientID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	response := mapPatientToResponse(patient)

	c.JSON(http.StatusOK, response)
}
