package db

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var databaseString = "host=localhost user=neich password=12345 dbname=test-go-crud port=8001"

var DB *gorm.DB  

func DBConnection() {
	var error error

	DB, error = gorm.Open(postgres.Open(databaseString), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Conectado a la base de datos!")
	}
 }