package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"zadanie/internal/app"
	"zadanie/internal/config"
	"zadanie/internal/server"
	"zadanie/internal/server/router"
	"zadanie/pkg/zaplogger"

	"github.com/oklog/run"
	"go.uber.org/zap"
)

var ErrOsSignal = errors.New("got os signal")

func main() {

	mode := os.Getenv("MYAPP_MODE")

	config, err := config.ViperConfigurationProvider(mode, false)
	if err != nil {
		log.Fatal("Read config error: ", err)
	}

	logger, zapLoggerCleanup, err := zaplogger.Provider(mode)
	if err != nil {
		log.Fatal("zap logger provider: ", err)
	}
	defer zapLoggerCleanup()

	application := app.NewApp(storage)

	handler := router.NewHandler(application)
	router := router.NewRouter(handler, mode)
	server := server.NewServer(config.Server.Port, logger, router)

	var (
		serviceGroup        run.Group
		interruptionChannel = make(chan os.Signal, 1)
	)

	serviceGroup.Add(func() error {

		signal.Notify(interruptionChannel, syscall.SIGINT, syscall.SIGTERM)
		osSignal := <-interruptionChannel

		logger.Info("application", zap.Any("Got OS signal", osSignal))

		return fmt.Errorf("%w: %s", ErrOsSignal, osSignal)

	}, func(error) {
		interruptionChannel <- syscall.SIGINT

	})

}
