package repositories

import (
	"property_service/internal/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type RoomTypeRepositoryInterface interface {
	GetBaseRepo() Repository[entities.RoomType]
}

type roomTypeRepository struct {
	base Repository[entities.RoomType]
}

func NewRoomTypeRepository(collection *mongo.Collection) RoomTypeRepositoryInterface {
	return &roomTypeRepository{
		base: NewRepository[entities.RoomType](collection),
	}
}

func (r *roomTypeRepository) GetBaseRepo() Repository[entities.RoomType] {
	return r.base
}
