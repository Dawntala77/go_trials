package database

import (
	"fmt"
	"os"
	"sync"

	"example.com/myproject/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var once sync.Once

func ConnectT() (*gorm.DB, error) {
	// use existing connection if already exists
	if db != nil {
		fmt.Println("DB connection exists")
		return db, nil
	}
	fmt.Println("No db connection, connecting...")

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DATABASE")
	port := os.Getenv("DATABASE_PORT")

	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Printf("\nSuccessfully connected to database!\n")

	//--- Sync Models ---//
	err = db.AutoMigrate(&models.Lego{})
	if err != nil {
		fmt.Println("Failed to migrate lego model")
		return nil, err
	}

	err = db.AutoMigrate(&models.Members{})
	if err != nil {
		fmt.Println("Failed to migrate members model")
		return nil, err
	}
	fmt.Printf("\nModels synchronized successfully!\n")

	return db, nil
}
