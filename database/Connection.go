package database

import (
	"fmt"
	"log"

	"github.com/HeyAditya95/task-manager-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB //global variabel to store database connection

// connection to connect to postgress
func ConnectDB() {
	var err error
	//DSN (Data Source Name) →is this  Connection details for PostgreSQL
	dsn := "host=localhost user=adityajha101 password=123456789 dbname=task_manager port=5432 sslmode=disable"

	//open the communication
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database : ", err)
	}

	//assign db to the global variable
	DB = db
	fmt.Println("✅Database connected successfull !")
	DB.AutoMigrate(&models.Task{})
}
