package repositories

import (
	"sender_service/internal/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepositoryInterface interface {
	GetBaseRepo() *Repository[entities.Category]
}

type CategoryRepository struct {
	base *Repository[entities.Category]
}

func NewCategoryRepository(collection *mongo.Collection) CategoryRepositoryInterface {
	return &CategoryRepository{
		base: &Repository[entities.Category]{
			collection: collection,
		},
	}
}

func (r *CategoryRepository) GetBaseRepo() *Repository[entities.Category] {
	return r.base
}
