package db

import (
	"http_server/internal/calculationService"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error

	//connect to Db using credentials
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Cannot connect to db: %v", err)
	}
	//creating model MyCalculation in the DB
	if err := db.AutoMigrate(&calculationService.MyCalculation{}); err != nil {
		log.Fatalf("Cannot migrate db: %v", err)
	}
	return db, nil
}
