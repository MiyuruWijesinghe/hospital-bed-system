package controllers

import (
	"net/http"
	"time"

	"hospital/config"
	"hospital/models"

	"github.com/gin-gonic/gin"
)

// Admit Patient & Assign Bed
func AdmitPatient(c *gin.Context) {

	var input struct {
		PatientID uint `json:"patientId"`
		WardID    uint `json:"wardId"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check patient exists
	var patient models.Patient
	if err := config.DB.First(&patient, input.PatientID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	// Find first available bed in that ward
	var bed models.Bed

	err := config.DB.
		Joins("JOIN rooms ON rooms.id = beds.room_id").
		Where("rooms.ward_id = ?", input.WardID).
		Where("beds.status = ?", "AVAILABLE").
		First(&bed).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No available beds in this ward"})
		return
	}

	// Create admission
	admission := models.Admission{
		PatientID: input.PatientID,
		BedID:     bed.ID,
		Status:    "ACTIVE",
	}

	config.DB.Create(&admission)

	// Update bed status to OCCUPIED
	bed.Status = "OCCUPIED"
	config.DB.Save(&bed)

	c.JSON(http.StatusOK, gin.H{
		"message":   "Patient admitted successfully",
		"admission": admission,
	})
}

// Discharge Patient
func DischargePatient(c *gin.Context) {

	admit_id := c.Param("admit_id")

	// Find admission
	var admission models.Admission
	if err := config.DB.First(&admission, admit_id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admission not found"})
		return
	}

	// Only active admissions can be discharged
	if admission.Status != "ACTIVE" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Admission already closed"})
		return
	}

	// Update admission
	now := time.Now()
	admission.Status = "DISCHARGED"
	admission.DischargedAt = &now

	config.DB.Save(&admission)

	// Free bed
	var bed models.Bed
	config.DB.First(&bed, admission.BedID)

	bed.Status = "CLEANING"
	config.DB.Save(&bed)

	c.JSON(http.StatusOK, gin.H{
		"message": "Patient discharged successfully",
	})
}

func GetAdmissions(c *gin.Context) {
	var admissions []models.Admission
	config.DB.Find(&admissions)

	c.JSON(http.StatusOK, admissions)
}

func GetAdmissionByID(c *gin.Context) {
	admissionID := c.Param("admission_id")

	var admission models.Admission
	if err := config.DB.First(&admission, admissionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admission details not found"})
		return
	}

	c.JSON(http.StatusOK, admission)
}
