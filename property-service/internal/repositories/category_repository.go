package repositories

import (
	"property_service/internal/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepositoryInterface interface {
	GetBaseRepo() Repository[entities.Category]
}

type categoryRepository struct {
	base Repository[entities.Category]
}

func NewCategoryRepository(collection *mongo.Collection) CategoryRepositoryInterface {
	return &categoryRepository{
		base: NewRepository[entities.Category](collection),
	}
}

func (r *categoryRepository) GetBaseRepo() Repository[entities.Category] {
	return r.base
}
