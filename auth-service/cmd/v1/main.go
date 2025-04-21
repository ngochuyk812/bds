package main

import (
	"auth_service/interfaces/v1/connectrpc"
	"auth_service/internal/infra"
	"os"
	"os/signal"

	"auth_service/internal/repository"
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	bus "auth_service/internal/infra/bus"

	infrastructurecore "github.com/ngochuyk812/building_block/infrastructure/core"
	"github.com/ngochuyk812/building_block/infrastructure/databases"
	"github.com/ngochuyk812/building_block/pkg/config"
)

var (
	brokers = os.Getenv("BROKERS_EVENTBUS")
	topic   = os.Getenv("TOPIC_EVENTBUS")
)

func main() {
	policiesPath := &map[string][]string{
		"/greet.v1.GreetService/Greet": {"user"},
	}
	config := config.NewConfigEnv()
	config.PoliciesPath = policiesPath
	infa := infrastructurecore.NewInfra(config)
	infa.InjectSQL(databases.MYSQL)
	infa.InjectCache(config.RedisConnect, config.RedisPass)
	unf := repository.NewUnitOfWork(infa.GetDatabase().GetWriteDB(), infa.GetDatabase().GetReadDB())
	cabin := infra.NewCabin(infa, unf)
	bus.InjectBus(cabin)
	// infa.InjectEventbus(brokers, topic)

	app := infrastructurecore.NewServe(":"+config.Port, infa.GetLogger())
	path, handler := connectrpc.NewAuthServer(cabin)
	app.Mux.Handle(path, handler)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go app.Run()
	<-c
	fmt.Println("shutting down...")

}
