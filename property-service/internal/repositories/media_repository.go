package repositories

import (
	"property_service/internal/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type MediaRepositoryInterface interface {
	GetBaseRepo() *Repository[entities.Media]
}

type mediaRepository struct {
	base *Repository[entities.Media]
}

func NewMediaRepository(collection *mongo.Collection) MediaRepositoryInterface {
	return &mediaRepository{
		base: &Repository[entities.Media]{
			collection: collection,
		},
	}
}

func (r *mediaRepository) GetBaseRepo() *Repository[entities.Media] {
	return r.base
}
