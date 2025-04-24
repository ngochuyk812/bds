package infra

import (
	infrastructurecore "github.com/ngochuyk812/building_block/infrastructure/core"
)

type Cabin interface {
	GetInfra() infrastructurecore.IInfra
}
type cabin struct {
	infra infrastructurecore.IInfra
}

func NewCabin(infra infrastructurecore.IInfra) Cabin {
	cabin := &cabin{
		infra: infra,
	}
	return cabin
}

func (c *cabin) GetInfra() infrastructurecore.IInfra {
	return c.infra
}
