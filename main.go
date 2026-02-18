package main

import (
	"hospital/config"
	"hospital/models"
	"hospital/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()

	err := config.DB.AutoMigrate(
		&models.Ward{},
		&models.Room{},
		&models.Bed{},
		&models.Patient{},
		&models.Admission{},
	)

	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}

	log.Println("✅ Database migrated successfully")

	routes.SetupRoutes(r)
	log.Println("🚀 Server running on http://localhost:8080")
	r.Run(":8080")
}
