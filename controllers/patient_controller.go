package controllers

import (
	"hospital/config"
	"hospital/models"
	"hospital/requests"
	"hospital/responses"
	"hospital/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

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

	response := utils.MapPatientToResponse(patient)

	c.JSON(http.StatusCreated, response)
}

func GetPatients(c *gin.Context) {
	var patients []models.Patient
	config.DB.Find(&patients)

	var response []responses.PatientResponse

	for _, p := range patients {
		response = append(response, utils.MapPatientToResponse(p))
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

	response := utils.MapPatientToResponse(patient)

	c.JSON(http.StatusOK, response)
}

func UpdatePatient(c *gin.Context) {

	id := c.Param("patient_id")

	var patient models.Patient

	if err := config.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	// Bind input
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

	// Update fields
	patient.FirstName = req.FirstName
	patient.LastName = req.LastName
	patient.Gender = req.Gender
	patient.DOB = dob

	config.DB.Save(&patient)

	response := utils.MapPatientToResponse(patient)

	c.JSON(http.StatusOK, response)
}

func DeletePatient(c *gin.Context) {
	id := c.Param("patient_id")

	var patient models.Patient

	if err := config.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	config.DB.Delete(&patient)

	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted successfully"})
}
