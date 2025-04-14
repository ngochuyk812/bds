package infra

import (
	"auth_service/internal/repository"

	infrastructurecore "github.com/ngochuyk812/building_block/infrastructure/core"
)

type Cabin interface {
	GetInfra() infrastructurecore.IInfra
	GetUnitOfWork() repository.UnitOfWork
}
type cabin struct {
	infra      infrastructurecore.IInfra
	unitOfWork repository.UnitOfWork
}

func NewCabin(infra infrastructurecore.IInfra, unitOfWork repository.UnitOfWork) Cabin {
	cabin := &cabin{
		infra:      infra,
		unitOfWork: unitOfWork,
	}
	return cabin
}

func (c *cabin) GetInfra() infrastructurecore.IInfra {
	return c.infra
}

func (c *cabin) GetUnitOfWork() repository.UnitOfWork {
	return c.unitOfWork
}
