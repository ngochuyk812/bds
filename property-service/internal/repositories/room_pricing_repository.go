package repositories

import (
	"sender_service/internal/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type RoomPricingRepositoryInterface interface {
	GetBaseRepo() *Repository[entities.RoomPricing]
}

type RoomPricingRepository struct {
	base *Repository[entities.RoomPricing]
}

func NewRoomPricingRepository(collection *mongo.Collection) RoomPricingRepositoryInterface {
	return &RoomPricingRepository{
		base: &Repository[entities.RoomPricing]{
			collection: collection,
		},
	}
}

func (r *RoomPricingRepository) GetBaseRepo() *Repository[entities.RoomPricing] {
	return r.base
}
