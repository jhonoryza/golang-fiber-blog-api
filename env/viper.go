package env

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

var env *viper.Viper

func LoadEnv() {
	env = viper.New()
	env.SetConfigFile(".env")
	if err := env.ReadInConfig(); err != nil {
		log.Error(err)
	}
}

func GetEnv() *viper.Viper {
	if env == nil {
		LoadEnv()
	}
	return env
}
