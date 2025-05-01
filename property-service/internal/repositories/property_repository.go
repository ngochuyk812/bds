package repositories

import (
	"sender_service/internal/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type PropertyRepositoryInterface interface {
	GetBaseRepo() *Repository[entities.Property]
}

type PropertyRepository struct {
	base *Repository[entities.Property]
}

func NewPropertyRepository(collection *mongo.Collection) PropertyRepositoryInterface {
	return &PropertyRepository{
		base: &Repository[entities.Property]{
			collection: collection,
		},
	}
}

func (r *PropertyRepository) GetBaseRepo() *Repository[entities.Property] {
	return r.base
}
