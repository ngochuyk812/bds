package repositories

import (
	"sender_service/internal/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type AmenityRepositoryInterface interface {
	GetBaseRepo() *Repository[entities.Amenity]
}

type AmenityRepository struct {
	base *Repository[entities.Amenity]
}

func NewAmenityRepository(collection *mongo.Collection) AmenityRepositoryInterface {
	return &AmenityRepository{
		base: &Repository[entities.Amenity]{
			collection: collection,
		},
	}
}

func (r *AmenityRepository) GetBaseRepo() *Repository[entities.Amenity] {
	return r.base
}
