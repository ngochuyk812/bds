package repositories

import (
	"property_service/internal/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type PropertyTypeRepositoryInterface interface {
	GetBaseRepo() Repository[entities.PropertyType]
}

type propertyTypeRepository struct {
	base Repository[entities.PropertyType]
}

func NewPropertyTypeRepository(collection *mongo.Collection) PropertyTypeRepositoryInterface {
	return &propertyTypeRepository{
		base: NewRepository[entities.PropertyType](collection),
	}
}

func (r *propertyTypeRepository) GetBaseRepo() Repository[entities.PropertyType] {
	return r.base
}
