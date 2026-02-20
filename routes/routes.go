package routes

import (
	"hospital/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	hospital := r.Group("/hospital")
	{
		// Wards
		wards := hospital.Group("/wards")
		{
			wards.GET("/", controllers.GetWards)
			wards.GET("/:ward_id", controllers.GetWardByID)
			wards.POST("/", controllers.CreateWard)
		}

		// Rooms
		rooms := hospital.Group("/rooms")
		{
			rooms.GET("/", controllers.GetRooms)
			rooms.GET("/:room_id", controllers.GetRoomByID)
			rooms.GET("/ward/:ward_id", controllers.GetRoomsByWard)
			rooms.POST("/", controllers.CreateRoom)
		}

		// Beds
		beds := hospital.Group("/beds")
		{
			beds.GET("/", controllers.GetBeds)
			beds.GET("/:bed_id", controllers.GetBedByID)
			beds.GET("/room/:room_id", controllers.GetBedsByRoom)
			beds.POST("/", controllers.CreateBed)
			beds.PUT("/:bed_id/status", controllers.UpdateBedStatus)
		}

		// Patients
		patients := hospital.Group("/patients")
		{
			patients.GET("/", controllers.GetPatients)
			patients.GET("/:patient_id", controllers.GetPatientByID)
			patients.POST("/", controllers.CreatePatient)
		}

	}
}
