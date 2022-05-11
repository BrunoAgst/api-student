package database

import (
	"api-rest/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectDatabase() {
	conectionString := "host=localhost user=admin password=layla123 dbname=postgres port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(conectionString))
	if err != nil {
		log.Panic("Error dont connect to database:", err)
	}

	DB.AutoMigrate(&models.Student{})

}
