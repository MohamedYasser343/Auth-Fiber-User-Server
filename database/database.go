package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error

	dsn := "host=127.0.0.1 port=5432 user=mohamed password=momo23 dbname=Learn_Fiber"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed To Connect To Database")
	}

	log.Println("Connected To Database")
}
