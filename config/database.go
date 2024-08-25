package config

import (
	"fiber_blog/env"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDatabase() *gorm.DB {
	logMode := logger.Info
	if env.GetEnv().GetString("APP_ENV") != "local" {
		logMode = logger.Silent
	}

	db, err := gorm.Open(postgres.Open(env.GetEnv().GetString("DATABASE_URL")), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logMode),
	})
	if err != nil {
		fmt.Println("failed to connect database")
	}
	fmt.Println("connected to database")
	return db
}
