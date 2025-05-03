package repositories

import (
	"property_service/internal/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type RoomPricingRepositoryInterface interface {
	GetBaseRepo() Repository[entities.RoomPricing]
}

type roomPricingRepository struct {
	base Repository[entities.RoomPricing]
}

func NewRoomPricingRepository(collection *mongo.Collection) RoomPricingRepositoryInterface {
	return &roomPricingRepository{
		base: NewRepository[entities.RoomPricing](collection),
	}
}

func (r *roomPricingRepository) GetBaseRepo() Repository[entities.RoomPricing] {
	return r.base
}
