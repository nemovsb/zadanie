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
	"zadanie/internal/storage/storage_mock"
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

	storage := storage_mock.NewStorageMock()

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

	serviceGroup.Add(func() error {

		logger.Info("application", zap.String("event", "Application logic started"))
		return application.Run()

	}, func(error) {

		err = application.Shutdown()
		logger.Info("application", zap.String("event", "Application logic shutdown"))
	})

	serviceGroup.Add(func() error {

		logger.Info("application", zap.String("event", "Http API started"))
		return server.Run()

	}, func(error) {

		err = server.Shutdown()
		logger.Info("application", zap.String("event", "Http API shutdown"))
	})

	err = serviceGroup.Run()
	logger.Info("services stopped", zap.Error(err))

}
