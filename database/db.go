package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	err error
)

func StablishConnection() {
	connectionInfo := "host=localhost user=golang password=golang dbname=postgresql port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionInfo))

	if err !=  nil {
		log.Panic("Error connecting to database!")
	}
}
