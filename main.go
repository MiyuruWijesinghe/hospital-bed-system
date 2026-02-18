package main

import (
	"hospital/config"
	"hospital/models"
	"log"
)

func main() {
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
}
