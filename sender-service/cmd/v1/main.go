package main

import (
	"os"
	events "sender_service/internal/app/event_handlers"
	"sender_service/internal/infra"
	"sender_service/internal/infra/bus"

	"fmt"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	infrastructurecore "github.com/ngochuyk812/building_block/infrastructure/core"
	"github.com/ngochuyk812/building_block/infrastructure/eventbus"
	"github.com/ngochuyk812/building_block/infrastructure/eventbus/kafka"
	"github.com/ngochuyk812/building_block/pkg/config"
)

func main() {
	var (
		brokers = os.Getenv("BROKERS_EVENTBUS")
		topic   = os.Getenv("TOPIC_EVENTBUS")
		group   = os.Getenv("GROUP_ID_EVENTBUS")
	)
	policiesPath := &map[string][]string{}
	config := config.NewConfigEnv()
	config.PoliciesPath = policiesPath
	inra := infrastructurecore.NewInfra(config)
	cabin := infra.NewCabin(inra)
	bus.InjectBus(cabin)

	consumer, err := kafka.NewConsumer(brokers, topic, group)
	if err != nil {
		panic(fmt.Sprintf("cannot connect consumer: %s", err.Error()))
	}

	registerHandlerEventbus(consumer, cabin)
	go consumer.Run()

}

func registerHandlerEventbus(consumer eventbus.Consumer, cabin infra.Cabin) {
	consumer.RegisterHandler(events.NewIntegrationEventSendMailHandler(cabin))
}
