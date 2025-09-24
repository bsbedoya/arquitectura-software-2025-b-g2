package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect recibe el DSN como parámetro
func Connect(dsn string) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB connect error: %v", err)
	}

	DB = db
	log.Println("✅ Connected to MySQL")
}
