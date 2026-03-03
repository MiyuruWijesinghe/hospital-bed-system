package controllers

import (
	"hospital/config"
	"hospital/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateWard(c *gin.Context) {
	var ward models.Ward

	if err := c.ShouldBindJSON(&ward); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&ward)
	c.JSON(http.StatusCreated, ward)
}

func GetWards(c *gin.Context) {
	var wards []models.Ward
	config.DB.Find(&wards)

	c.JSON(http.StatusOK, wards)
}

func GetWardByID(c *gin.Context) {
	wardID := c.Param("ward_id")

	var ward models.Ward
	if err := config.DB.First(&ward, wardID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ward not found"})
		return
	}

	c.JSON(http.StatusOK, ward)
}

func UpdateWard(c *gin.Context) {
	id := c.Param("ward_id")

	var ward models.Ward

	if err := config.DB.First(&ward, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ward not found"})
		return
	}

	// Bind input
	var input models.Ward
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	ward.Name = input.Name
	ward.Type = input.Type
	ward.Floor = input.Floor

	config.DB.Save(&ward)

	c.JSON(http.StatusOK, ward)
}

func DeleteWard(c *gin.Context) {
	id := c.Param("ward_id")

	var ward models.Ward

	if err := config.DB.First(&ward, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ward not found"})
		return
	}

	config.DB.Delete(&ward)

	c.JSON(http.StatusOK, gin.H{"message": "Ward deleted successfully"})
}
