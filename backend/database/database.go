package database

import (
	"fmt"
	"taskido/backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsn := "upsaurav12:SA@2003up_@tcp(127.0.0.1:3306)/taskido?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connected Database Successfully")

	err = database.AutoMigrate(&models.Task{})

	if err != nil {
		panic("Failed to Migrate Database")
	}

	fmt.Println("Database Migrated Successfully")

	DB = database

}
