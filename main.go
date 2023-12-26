package main

import (
	"log"
	"os"
	"zadanie/internal/config"
	"zadanie/pkg/zaplogger"
)

func main() {

	config, err := config.ViperConfigurationProvider(os.Getenv("MYAPP_MODE"), false)
	if err != nil {
		log.Fatal("Read config error: ", err)
	}

	logger, zapLoggerCleanup, err := zaplogger.Provider(os.Getenv("MYAPP_MODE"))
	if err != nil {
		log.Fatal("zap logger provider: ", err)
	}
	defer zapLoggerCleanup()

}
