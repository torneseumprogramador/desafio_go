package database

import (
	"http_fiber_api/pkg/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	dsn := config.GetDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	return db
}
