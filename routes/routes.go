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
			wards.GET("", controllers.GetWards)
			wards.GET("/:ward_id", controllers.GetWardByID)
			wards.POST("", controllers.CreateWard)
			wards.PUT("/:ward_id", controllers.UpdateWard)
			wards.DELETE("/:ward_id", controllers.DeleteWard)
		}

		// Rooms
		rooms := hospital.Group("/rooms")
		{
			rooms.GET("", controllers.GetRooms)
			rooms.GET("/:room_id", controllers.GetRoomByID)
			rooms.GET("/ward/:ward_id", controllers.GetRoomsByWard)
			rooms.POST("", controllers.CreateRoom)
			rooms.PUT("/:room_id", controllers.UpdateRoom)
			rooms.DELETE("/:room_id", controllers.DeleteRoom)
		}

		// Beds
		beds := hospital.Group("/beds")
		{
			beds.GET("", controllers.GetBeds)
			beds.GET("/:bed_id", controllers.GetBedByID)
			beds.GET("/room/:room_id", controllers.GetBedsByRoom)
			beds.POST("", controllers.CreateBed)
			beds.PUT("/:bed_id/status", controllers.UpdateBedStatus)
			beds.PUT("/:bed_id", controllers.UpdateBed)
			beds.DELETE("/:bed_id", controllers.DeleteBed)
		}

		// Patients
		patients := hospital.Group("/patients")
		{
			patients.GET("", controllers.GetPatients)
			patients.GET("/:patient_id", controllers.GetPatientByID)
			patients.POST("", controllers.CreatePatient)
		}

		// Admissions
		admissions := hospital.Group("/admissions")
		{
			admissions.GET("", controllers.GetAdmissions)
			admissions.GET("/:admission_id", controllers.GetAdmissionByID)
			admissions.POST("/admit", controllers.AdmitPatient)
			admissions.PUT("/discharge/:admit_id", controllers.DischargePatient)
		}

		// Dashboard
		dashboard := hospital.Group("/dashboard")
		{
			dashboard.GET("", controllers.GetDashboard)
		}
	}
}
