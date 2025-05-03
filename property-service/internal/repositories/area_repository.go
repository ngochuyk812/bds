package repositories

import (
	"property_service/internal/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type AreaRepositoryInterface interface {
	GetBaseRepo() *Repository[entities.Area]
}

type areaRepository struct {
	base *Repository[entities.Area]
}

func NewAreaRepository(collection *mongo.Collection) AreaRepositoryInterface {
	return &areaRepository{
		base: &Repository[entities.Area]{
			collection: collection,
		},
	}
}

func (r *areaRepository) GetBaseRepo() *Repository[entities.Area] {
	return r.base
}
