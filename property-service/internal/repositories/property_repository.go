package repositories

import (
	"property_service/internal/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type PropertyRepositoryInterface interface {
	GetBaseRepo() Repository[entities.Property]
}

type propertyRepository struct {
	base Repository[entities.Property]
}

func NewPropertyRepository(collection *mongo.Collection) PropertyRepositoryInterface {
	return &propertyRepository{
		base: NewRepository[entities.Property](collection),
	}
}

func (r *propertyRepository) GetBaseRepo() Repository[entities.Property] {
	return r.base
}
