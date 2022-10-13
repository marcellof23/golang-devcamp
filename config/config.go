package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Setup : initializing mysql database
func Init() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("host")
	port := os.Getenv("port")
	user := os.Getenv("user")
	dbname := os.Getenv("dbname")
	password := os.Getenv("password")
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
