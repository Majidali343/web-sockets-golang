package dbconnect

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func Dbconnection() *gorm.DB {
	var db *gorm.DB
	var err error

	host := "localhost"
	username := "postgres"
	password := "Majid"
	dbName := "chat"

	// Construct the connection string
	connectionString := fmt.Sprintf("host=%s user=%s dbname=%s password=%s   sslmode=disable",
		host, username, dbName, password)

	// Connect to PostgreSQL
	db, err = gorm.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)

	}

	return db
}
