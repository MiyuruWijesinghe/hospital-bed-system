package routes

import (
	"hospital/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	api := r.Group("/api")
	{
		api.GET("/beds", controllers.GetBeds)
		api.POST("/beds", controllers.CreateBed)
		api.PUT("/beds/:id/status", controllers.UpdateBedStatus)
	}
}
