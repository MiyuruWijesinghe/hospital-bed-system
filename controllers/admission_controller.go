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
		PatientID uint   `json:"patientId"`
		BedID     uint   `json:"bedId"`
		Reason    string `json:"reason"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := config.DB.Begin() // start transaction

	// Check patient exists
	var patient models.Patient
	if err := tx.First(&patient, input.PatientID).Error; err != nil {
		tx.Rollback()
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

	// check bed
	var bed models.Bed
	if err := tx.First(&bed, input.BedID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Bed not found"})
		return
	}

	if bed.Status != "AVAILABLE" {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bed is not available"})
		return
	}

	// Create admission
	admission := models.Admission{
		PatientID: input.PatientID,
		BedID:     input.BedID,
		Reason:    input.Reason,
		Status:    "ACTIVE",
	}

	if err := tx.Create(&admission).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create admission"})
		return
	}

	// Update bed status
	if err := tx.Model(&bed).Update("status", "OCCUPIED").Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update bed"})
		return
	}

	tx.Commit() // commit transaction

	c.JSON(http.StatusCreated, gin.H{
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

	// Check if already discharged
	if admission.Status != "ACTIVE" {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Patient already discharged"})
		return
	}

	// Update admission
	now := time.Now()

	admission.Status = "DISCHARGED"
	admission.DischargedAt = &now

	if err := tx.Save(&admission).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update admission"})
		return
	}

	// Find bed
	var bed models.Bed
	if err := tx.First(&bed, admission.BedID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Bed not found"})
		return
	}

	// Set bed to CLEANING
	if err := tx.Model(&bed).Update("status", "CLEANING").Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update bed"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"message":   "Patient discharged successfully",
		"bedStatus": "CLEANING",
	})
}

func GetAdmissions(c *gin.Context) {

	var admissions []models.Admission

	config.DB.
		Preload("Patient").
		Preload("Bed").
		Preload("Bed.Room").
		Preload("Bed.Room.Ward").
		Find(&admissions)

	c.JSON(http.StatusOK, admissions)
}

func GetAdmissionByID(c *gin.Context) {

	admissionID := c.Param("admission_id")

	var admission models.Admission

	err := config.DB.
		Preload("Patient").
		Preload("Bed").
		Preload("Bed.Room").
		Preload("Bed.Room.Ward").
		First(&admission, admissionID).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admission details not found"})
		return
	}

	c.JSON(http.StatusOK, admission)
}
