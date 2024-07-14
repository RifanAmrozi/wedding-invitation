package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	DB = db
	fmt.Println("Database connected")
}
