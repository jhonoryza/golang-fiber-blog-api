package app

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

func GetConfig() *viper.Viper {
	config := viper.New()
	config.SetConfigFile(".env")
	if err := config.ReadInConfig(); err != nil {
		log.Error(err)
	}
	return config
}
