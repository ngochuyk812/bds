package main

import (
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/ngochuyk812/building_block/pkg/config"
)

func main() {

	policiesPath := &map[string][]string{}
	config := config.NewConfigEnv()
	config.PoliciesPath = policiesPath

	select {}

}
