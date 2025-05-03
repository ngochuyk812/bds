package repositories

import (
	"property_service/internal/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type AreaRepositoryInterface interface {
	GetBaseRepo() Repository[entities.Area]
}

type areaRepository struct {
	base Repository[entities.Area]
}

func NewAreaRepository(collection *mongo.Collection) AreaRepositoryInterface {
	return &areaRepository{
		base: NewRepository[entities.Area](collection),
	}
}

func (r *areaRepository) GetBaseRepo() Repository[entities.Area] {
	return r.base
}
