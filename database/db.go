package database

import (
	"example.com/myproject/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sync"
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

	connectionString := os.Getenv("DB_URL")

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
