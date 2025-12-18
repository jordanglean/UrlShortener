package db

import (
	"github.com/bytedance/gopkg/util/logger"
	"github.com/jordanglean/UrlShortener/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		logger.Error("Error connecting to database: ", err)
		return
	}

	err = db.AutoMigrate(&models.ShortenURL{})
	if err != nil {
		logger.Error("Error creating table migration: ", err)
		return
	}

}
