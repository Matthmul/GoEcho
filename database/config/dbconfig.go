package database

import (
	"fmt"

	"store/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB = nil

func Init() {
	db, err := gorm.Open(sqlite.Open("./database/store.db"), &gorm.Config{})
	if err != nil {
		fmt.Printf("storage error")
	}

	db.AutoMigrate(
		&models.Cart{},
		&models.Category{},
		&models.Order{},
		&models.Payment{},
		&models.Product{},
	)

	DB = db
	fmt.Printf("DB connected")
}