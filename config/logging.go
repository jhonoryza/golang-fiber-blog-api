package config

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"os"
)

func InitLogging() {
	appLog, err := os.OpenFile("app.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(appLog)
	fmt.Println("log initialized")
}
