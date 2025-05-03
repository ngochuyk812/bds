package repositories

import (
	"property_service/internal/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type AmenityRepositoryInterface interface {
	GetBaseRepo() *Repository[entities.Amenity]
}

type amenityRepository struct {
	base *Repository[entities.Amenity]
}

func NewAmenityRepository(collection *mongo.Collection) AmenityRepositoryInterface {
	return &amenityRepository{
		base: &Repository[entities.Amenity]{
			collection: collection,
		},
	}
}

func (r *amenityRepository) GetBaseRepo() *Repository[entities.Amenity] {
	return r.base
}
