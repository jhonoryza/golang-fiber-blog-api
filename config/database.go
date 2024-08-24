package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDatabase() {
	logMode := logger.Info
	if GetEnv().GetString("APP_ENV") != "local" {
		logMode = logger.Silent
	}

	newDb, err := gorm.Open(postgres.Open(GetEnv().GetString("DATABASE_URL")), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logMode),
	})
	if err != nil {
		fmt.Println("failed to connect database")
	}
	fmt.Println("connected to database")
	db = newDb
}

func GetDB() *gorm.DB {
	if db == nil {
		InitDatabase()
	}
	return db
}
