package repositories

import (
	"sender_service/internal/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type PropertyTypeRepositoryInterface interface {
	GetBaseRepo() *Repository[entities.PropertyType]
}

type PropertyTypeRepository struct {
	base *Repository[entities.PropertyType]
}

func NewPropertyTypeRepository(collection *mongo.Collection) PropertyTypeRepositoryInterface {
	return &PropertyTypeRepository{
		base: &Repository[entities.PropertyType]{
			collection: collection,
		},
	}
}

func (r *PropertyTypeRepository) GetBaseRepo() *Repository[entities.PropertyType] {
	return r.base
}
