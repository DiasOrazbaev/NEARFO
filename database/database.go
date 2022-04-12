package database

import (
	"log"
	"os"

	"github.com/DiasOrazbaev/NEARFO/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDB() *gorm.DB {
	dsn, envExists := os.LookupEnv("DB_DSN")
	if !envExists {
		panic("Not DSN string")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not init database")
	}
	migrate(db)
	return db
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.Shop{}, &models.Good{})
	if err != nil {
		log.Print("Could not init database")
		return
	}
	log.Print("Database migrated")
}
