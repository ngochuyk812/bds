package main

import (
	"auth_service/interfaces/v1/connectrpc"
	"auth_service/internal/infra"
	"auth_service/internal/infra/database"
	"auth_service/internal/usecase"
	"os"
	"os/signal"

	"auth_service/internal/repository"
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	infrastructurecore "github.com/ngochuyk812/building_block/infrastructure/core"
	"github.com/ngochuyk812/building_block/pkg/config"
)

var (
	brokers   = os.Getenv("BROKERS_EVENTBUS")
	topic     = os.Getenv("TOPIC_EVENTBUS")
	dbConnect = os.Getenv("DB_CONNECTION")
	dbName    = os.Getenv("DB_NAME")
)

func main() {
	policiesPath := &map[string][]string{
		"/greet.v1.GreetService/Greet": {"user"},
	}
	config := config.NewConfigEnv()
	config.PoliciesPath = policiesPath
	infa := infrastructurecore.NewInfra(config)
	// infa.InjectCache(config.RedisConnect, config.RedisPass)
	// infa.InjectEventbus(brokers, topic)

	db := database.NewSQLDB(dbConnect, dbName)
	unf := repository.NewUnitOfWork(db)

	cabin := infra.NewCabin(infa, unf)
	useCases := usecase.NewUsecaseManager(cabin)

	app := infrastructurecore.NewServe(":"+config.Port, infa.GetLogger())
	path, handler := connectrpc.NewAuthServer(useCases, cabin)
	app.Mux.Handle(path, handler)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go app.Run()
	<-c
	fmt.Println("shutting down...")

}
