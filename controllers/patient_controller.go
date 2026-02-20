package controllers

import (
	"hospital/config"
	"hospital/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePatient(c *gin.Context) {
	var patient models.Patient

	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&patient)

	c.JSON(http.StatusOK, patient)
}

func GetPatients(c *gin.Context) {
	var patients []models.Patient
	config.DB.Find(&patients)

	c.JSON(http.StatusOK, patients)
}

func GetPatientByID(c *gin.Context) {
	patientID := c.Param("patient_id")

	var patient models.Patient
	if err := config.DB.First(&patient, patientID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	c.JSON(http.StatusOK, patient)
}
