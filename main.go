package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"zadanie/docs"
	"zadanie/internal/app"
	"zadanie/internal/config"
	"zadanie/internal/server"
	"zadanie/internal/server/router"
	"zadanie/internal/storage/pg"
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

	fmt.Printf("config = %+v\n", config)

	logger, zapLoggerCleanup, err := zaplogger.Provider(mode)
	if err != nil {
		log.Fatal("zap logger provider: ", err)
	}
	defer zapLoggerCleanup()

	docs.SetSwagger()

	//storage := storage_mock.NewStorageMock()
	storage, err := pg.NewPostgres(pg.NewConfig(
		config.Storage.UserName,
		config.Storage.Host,
		config.Storage.Port,
		config.Storage.Password,
		config.Storage.DBName,
	))
	if err != nil {
		log.Fatal("create storage error:", err)
	}

	application := app.NewApp(storage, logger)

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

	appCtx, cancel := context.WithCancel(context.Background())

	serviceGroup.Add(func() error {

		logger.Info("application", zap.String("event", "Application logic started"))
		return application.Run(appCtx)

	}, func(error) {

		err = application.Shutdown(cancel)
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
