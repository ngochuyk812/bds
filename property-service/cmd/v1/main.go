package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"property_service/interfaces/v1/connectrpc"
	"property_service/internal/config"
	"property_service/internal/infra"
	"property_service/internal/infra/db"
	"property_service/internal/usecases"
	"syscall"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	infrastructurecore "github.com/ngochuyk812/building_block/infrastructure/core"
	"github.com/ngochuyk812/building_block/infrastructure/logger"
)

func main() {
	ctx := context.Background()

	infrast := infrastructurecore.NewInfra()

	client := db.NewMongoClient(ctx, config.DbConnect)
	defer client.Disconnect(ctx)
	uow := db.NewUnitOfWork(client, config.DbName)

	cabin := infra.NewCabin(infrast, uow)
	usecases := usecases.NewUsecaseManager(cabin)

	mux := http.NewServeMux()

	path, handler := connectrpc.NewPropertyServer(cabin, usecases)
	mux.Handle(path, handler)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Port),
		Handler: mux,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		logger.Info(fmt.Sprintf("Starting server on: %s", config.Port))
		if err := server.ListenAndServe(); err != nil {
			logger.Error(fmt.Sprintf("Error starting server: %s", err))
		}
	}()

	<-c
	fmt.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error(fmt.Sprintf("Error shutting down server: %v", err))
	}

}
