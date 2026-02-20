package controllers

import (
	"hospital/config"
	"hospital/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBed(c *gin.Context) {
	var bed models.Bed

	if err := c.ShouldBindJSON(&bed); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&bed)
	c.JSON(http.StatusOK, bed)
}

func GetBeds(c *gin.Context) {
	var beds []models.Bed
	config.DB.Find(&beds)

	c.JSON(http.StatusOK, beds)
}

func GetBedByID(c *gin.Context) {
	bedID := c.Param("bed_id")

	var bed models.Bed
	if err := config.DB.First(&bed, bedID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bed not found"})
		return
	}

	c.JSON(http.StatusOK, bed)
}

func GetBedsByRoom(c *gin.Context) {
	roomID := c.Param("room_id")

	var beds []models.Bed
	config.DB.Where("room_id = ?", roomID).Find(&beds)

	c.JSON(http.StatusOK, beds)
}

func UpdateBedStatus(c *gin.Context) {
	bedID := c.Param("bed_id")

	var bed models.Bed
	if err := config.DB.First(&bed, bedID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bed not found"})
		return
	}

	var input struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bed.Status = input.Status
	config.DB.Save(&bed)

	c.JSON(http.StatusOK, bed)
}
