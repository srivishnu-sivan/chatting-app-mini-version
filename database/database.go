package database

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB


func ConnectDB(){
	dsn := "host=localhost user=postgres password=test@123 dbname=chat_app port=5432 sslmode=disable"
	// dsn := "host=localhost user=postgres password=test@123 dbname=chat_app port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err !=nil{
		fmt.Println("Error connecting to the database:", err)
		return
	}

	fmt.Println("Connected to the database successfully")

	DB = database
}