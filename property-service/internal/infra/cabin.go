package infra

import (
	"property_service/internal/infra/db"

	infrastructurecore "github.com/ngochuyk812/building_block/infrastructure/core"
)

type Cabin interface {
	GetInfra() infrastructurecore.IInfra
	GetUnitOfWork() db.UnitOfWork
}
type cabin struct {
	infra      infrastructurecore.IInfra
	unitOfWork db.UnitOfWork
}

func NewCabin(infra infrastructurecore.IInfra, unitOfWork db.UnitOfWork) Cabin {
	cabin := &cabin{
		infra:      infra,
		unitOfWork: unitOfWork,
	}
	return cabin
}

func (c *cabin) GetInfra() infrastructurecore.IInfra {
	return c.infra
}

func (c *cabin) GetUnitOfWork() db.UnitOfWork {
	return c.unitOfWork
}
