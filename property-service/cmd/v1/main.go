package main

import (
	"sender_service/internal/infra"
	"sender_service/internal/infra/bus"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	infrastructurecore "github.com/ngochuyk812/building_block/infrastructure/core"
	"github.com/ngochuyk812/building_block/pkg/config"
)

func main() {

	policiesPath := &map[string][]string{}
	config := config.NewConfigEnv()
	config.PoliciesPath = policiesPath
	inra := infrastructurecore.NewInfra(config)
	cabin := infra.NewCabin(inra)
	bus.InjectBus(cabin)

	select {}

}
