package dbconnect

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func Dbconnection() *gorm.DB {
	var db *gorm.DB
	var err error

	// godotenv.Load()
	// host := os.Getenv("DB_HOST")
	// username := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")

	host := "localhost"
	username := "postgres"
	password := "Majid"
	dbName := "postgres"

	// Construct the connection string
	connectionString := fmt.Sprintf("host=%s user=%s dbname=%s password=%s  host=postgres sslmode=disable",
		host, username, dbName, password)

	// Connect to PostgreSQL
	db, err = gorm.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)

	}

	return db
}
