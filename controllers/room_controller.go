package controllers

import (
	"hospital/config"
	"hospital/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRoom(c *gin.Context) {
	var room models.Room

	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check ward exists
	var ward models.Ward
	if err := config.DB.First(&ward, room.WardID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ward does not exist"})
		return
	}

	config.DB.Create(&room)
	c.JSON(http.StatusOK, room)
}

func GetRooms(c *gin.Context) {
	var rooms []models.Room
	config.DB.Find(&rooms)

	c.JSON(http.StatusOK, rooms)
}

func GetRoomByID(c *gin.Context) {
	roomID := c.Param("room_id")

	var room models.Room
	if err := config.DB.First(&room, roomID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	c.JSON(http.StatusOK, room)
}

func GetRoomsByWard(c *gin.Context) {
	wardID := c.Param("ward_id")

	var rooms []models.Room
	config.DB.Where("ward_id = ?", wardID).Find(&rooms)

	c.JSON(http.StatusOK, rooms)
}
