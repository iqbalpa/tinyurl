package config

import (
	"fmt"
	"log"
	"os"
	"tinyurl/models"

	"github.com/lpernett/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", 
							dbHost, 
							dbUsername, 
							dbPassword, 
							dbName, 
							dbPort)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
			log.Fatal("Failed to connect to database. ", err)
	}

	// migrate
	db.AutoMigrate(&models.ShortenedUrl{})

	return db
}