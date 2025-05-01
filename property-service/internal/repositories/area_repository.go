package repositories

import (
	"sender_service/internal/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type AreaRepositoryInterface interface {
	GetBaseRepo() *Repository[entities.Area]
}

type AreaRepository struct {
	base *Repository[entities.Area]
}

func NewAreaRepository(collection *mongo.Collection) AreaRepositoryInterface {
	return &AreaRepository{
		base: &Repository[entities.Area]{
			collection: collection,
		},
	}
}

func (r *AreaRepository) GetBaseRepo() *Repository[entities.Area] {
	return r.base
}
