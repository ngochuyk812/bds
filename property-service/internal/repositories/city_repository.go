package repositories

import (
	"sender_service/internal/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type CityRepositoryInterface interface {
	GetBaseRepo() *Repository[entities.City]
}

type CityRepository struct {
	base *Repository[entities.City]
}

func NewCityRepository(collection *mongo.Collection) CityRepositoryInterface {
	return &CityRepository{
		base: &Repository[entities.City]{
			collection: collection,
		},
	}
}

func (r *CityRepository) GetBaseRepo() *Repository[entities.City] {
	return r.base
}
