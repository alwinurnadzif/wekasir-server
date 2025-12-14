package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func DatabaseInit() error {
	DB_PORT := os.Getenv("database.port")
	DB_HOSTNAME := os.Getenv("database.hostname")
	DB_USERNAME := os.Getenv("database.username")
	DB_PASSWORD := os.Getenv("database.password")
	DB_NAME := os.Getenv("database.database")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		DB_HOSTNAME, DB_USERNAME, DB_PASSWORD, DB_NAME, DB_PORT)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		// Logger: logger.Discard,
	})
	if err != nil {
		return err
	}

	return nil
}
