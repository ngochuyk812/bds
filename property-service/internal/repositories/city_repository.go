package repositories

import (
	"property_service/internal/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type CityRepositoryInterface interface {
	GetBaseRepo() Repository[entities.City]
}

type cityRepository struct {
	base Repository[entities.City]
}

func NewCityRepository(collection *mongo.Collection) CityRepositoryInterface {
	return &cityRepository{
		base: NewRepository[entities.City](collection),
	}
}

func (r *cityRepository) GetBaseRepo() Repository[entities.City] {
	return r.base
}
