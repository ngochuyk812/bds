package repositories

import (
	"sender_service/internal/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type MediaRepositoryInterface interface {
	GetBaseRepo() *Repository[entities.Media]
}

type MediaRepository struct {
	base *Repository[entities.Media]
}

func NewMediaRepository(collection *mongo.Collection) MediaRepositoryInterface {
	return &MediaRepository{
		base: &Repository[entities.Media]{
			collection: collection,
		},
	}
}

func (r *MediaRepository) GetBaseRepo() *Repository[entities.Media] {
	return r.base
}
