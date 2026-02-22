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

	tx := config.DB.Begin() // start transaction

	// Check patient exists
	var patient models.Patient
	if err := config.DB.First(&patient, input.PatientID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	// Prevent duplicate active admission
	var existing models.Admission
	if err := tx.Where("patient_id = ? AND status = ?", input.PatientID, "ACTIVE").
		First(&existing).Error; err == nil {

		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Patient already admitted"})
		return
	}

	// Find first available bed in that ward
	var bed models.Bed

	err := tx.
		Joins("JOIN rooms ON rooms.id = beds.room_id").
		Where("rooms.ward_id = ?", input.WardID).
		Where("beds.status = ?", "AVAILABLE").
		Set("gorm:query_option", "FOR UPDATE"). // row lock
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

	if err := tx.Create(&admission).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create admission"})
		return
	}

	// Update bed status to OCCUPIED
	if err := tx.Model(&bed).Update("status", "OCCUPIED").Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update bed"})
		return
	}

	tx.Commit() // commit transaction

	c.JSON(http.StatusOK, gin.H{
		"message":   "Patient admitted successfully",
		"admission": admission,
	})
}

// Discharge Patient
func DischargePatient(c *gin.Context) {

	admit_id := c.Param("admit_id")

	tx := config.DB.Begin()

	// Find admission
	var admission models.Admission
	if err := tx.First(&admission, admit_id).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Admission not found"})
		return
	}

	// Only active admissions can be discharged
	if admission.Status != "ACTIVE" {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Already discharged"})
		return
	}

	// Update admission
	now := time.Now()
	admission.Status = "DISCHARGED"
	admission.DischargedAt = &now

	tx.Save(&admission)

	// Free bed
	var bed models.Bed
	tx.First(&bed, admission.BedID)
	tx.Model(&bed).Update("status", "CLEANING")

	tx.Commit()

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
