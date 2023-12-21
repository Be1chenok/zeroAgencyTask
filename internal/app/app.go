package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Be1chenok/zeroAgencyTask/internal/config"
	appHandler "github.com/Be1chenok/zeroAgencyTask/internal/delivery/handler"
	appServer "github.com/Be1chenok/zeroAgencyTask/internal/delivery/server"
	appLogger "github.com/Be1chenok/zeroAgencyTask/internal/logger"
	appRepository "github.com/Be1chenok/zeroAgencyTask/internal/repository"
	"github.com/Be1chenok/zeroAgencyTask/internal/repository/postgres"
	appService "github.com/Be1chenok/zeroAgencyTask/internal/service"
	"go.uber.org/zap"
)

func Run() {
	logger, err := appLogger.NewLogger()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Fatalf("failed to sync logger: %v", err)
		}
	}()
	appLog := logger.With(zap.String("component", "app"))

	conf, err := config.Init()
	if err != nil {
		appLog.Fatalf("failed to init config: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	postgres, err := postgres.New(conf, ctx)
	if err != nil {
		appLog.Fatalf("failed to connect database: %v", err)
	}
	cancel()

	repository := appRepository.New(logger, postgres)
	service := appService.New(repository, logger)
	handler := appHandler.New(conf, service)
	server := appServer.New(conf, *handler)
	server.InitRoutes()

	go func() {
		if err := server.Start(conf); err != nil {
			appLog.Fatalf("failed to start server: %v", err)
		}
	}()

	appLog.Infof("server is running on port %v", conf.Server.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGKILL)
	<-quit

	appLog.Infof("shuting down")

	if err := server.Shutdown(context.Background()); err != nil {
		appLog.Fatalf("failed to shut down server: %v", err)
	}

	if err := postgres.Close(); err != nil {
		appLog.Fatalf("failed to close database connection: %v", err)
	}
}
