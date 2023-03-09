package migration

import (
	"log"

	"github.com/MohamedYasser343/database"
	"github.com/MohamedYasser343/models"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Println(err)
	}

	log.Println("Database Migrated")
}
