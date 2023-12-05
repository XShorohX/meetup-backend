package db

import (
	"go-backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto Migrate Models
	database.AutoMigrate(&models.Speaker{}, &models.Report{}, &models.Meet{}, &models.QuizImage{})

	db = database
}

func GetDB() *gorm.DB {
	return db
}