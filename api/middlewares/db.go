package middlewares

// connecting to a PostgreSQL database with Go's database/sql package
import (
	"fmt"

	"example.com/myproject/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection() (*gorm.DB, error) {

	/*
	   variables required for connection string: connStr

	   user= (using default user for postgres database)
	   dbname= (using default database that comes with postgres)
	   password = (password used during initial setup)
	   host = (hostname or IP Address of server)
	   sslmode = (must be set to disabled unless using SSL)
	*/

	connStr := "user=postgres dbname=lego_builder password=dawn_tala host=localhost sslmode=disable"

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Lego{})
	db.AutoMigrate(&models.Members{})

	fmt.Printf("\nSuccessfully connected to database!\n")
	return db, err
}
