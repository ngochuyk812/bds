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
	"syscall"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	infrastructurecore "github.com/ngochuyk812/building_block/infrastructure/core"
	configinfra "github.com/ngochuyk812/building_block/pkg/config"
)

func main() {
	ctx := context.Background()

	config := config.NewConfig()
	configInfra := configinfra.NewConfigEnv()

	infrast := infrastructurecore.NewInfra(configInfra)

	client := db.NewMongoClient(ctx, config.DbConnection)
	defer client.Disconnect(ctx)
	uow := db.NewUnitOfWork(client, config.DBName)

	cabin := infra.NewCabin(infrast, uow)

	mux := http.NewServeMux()

	path, handler := connectrpc.NewPropertyServer(cabin)
	mux.Handle(path, handler)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", configInfra.Port),
		Handler: mux,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		infrast.GetLogger().Info(fmt.Sprintf("Starting server on: %s", configInfra.Port))
		if err := server.ListenAndServe(); err != nil {
			infrast.GetLogger().Error(fmt.Sprintf("Error starting server: %s", err))
		}
	}()

	<-c
	fmt.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		infrast.GetLogger().Error(fmt.Sprintf("Error shutting down server: %v", err))
	}

}
