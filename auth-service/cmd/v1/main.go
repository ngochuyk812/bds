package main

import (
	"auth_service/interfaces/v1/connectrpc"
	"auth_service/internal/infra"
	"auth_service/internal/infra/database"
	"auth_service/internal/usecase"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"auth_service/internal/repository"
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/ngochuyk812/building_block/infrastructure/logger"

	infrastructurecore "github.com/ngochuyk812/building_block/infrastructure/core"
	"github.com/ngochuyk812/building_block/pkg/config"
)

var (
	brokers   = os.Getenv("BROKERS_EVENTBUS")
	topic     = os.Getenv("TOPIC_EVENTBUS")
	dbConnect = os.Getenv("DB_CONNECTION")
	dbName    = os.Getenv("DB_NAME")
	port      = os.Getenv("SERVER_PORT")
)

func main() {
	policiesPath := &map[string][]string{
		"/greet.v1.GreetService/Greet": {"user"},
	}
	config := config.NewConfigEnv()
	config.PoliciesPath = policiesPath
	infrast := infrastructurecore.NewInfra()
	infrast.InjectCache(config.RedisConnect, config.RedisPass)
	// infrast.InjectEventbus(brokers, topic)

	db := database.NewSQLDB(dbConnect, dbName)
	unf := repository.NewUnitOfWork(db)

	cabin := infra.NewCabin(infrast, unf)
	useCases := usecase.NewUsecaseManager(cabin)

	path, handler := connectrpc.NewAuthServer(useCases, cabin)

	mux := http.NewServeMux()
	mux.Handle(path, handler)

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		logger.Info(fmt.Sprintf("Starting server on: %s", port))
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
