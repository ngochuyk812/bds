package main

import (
	"os"
	"os/signal"
	"sender_service/internal/infra"
	"sender_service/internal/infra/bus"
	"sender_service/internal/repository"

	"fmt"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	infrastructurecore "github.com/ngochuyk812/building_block/infrastructure/core"
	"github.com/ngochuyk812/building_block/pkg/config"
)

func main() {
	policiesPath := &map[string][]string{}
	config := config.NewConfigEnv()
	config.PoliciesPath = policiesPath
	infa := infrastructurecore.NewInfra(config)
	// infa.InjectSQL(databases.MYSQL)
	// infa.InjectCache(config.RedisConnect, config.RedisPass)
	unf := repository.NewUnitOfWork(infa.GetDatabase().GetWriteDB(), infa.GetDatabase().GetReadDB())
	cabin := infra.NewCabin(infa, unf)
	bus.InjectBus(cabin)
	app := infrastructurecore.NewServe(":"+config.Port, infa.GetLogger())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go app.Run()
	<-c
	fmt.Println("shutting down...")

}
