package repositories

import (
	"property_service/internal/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type RoomRepositoryInterface interface {
	GetBaseRepo() Repository[entities.Room]
}

type roomRepository struct {
	base Repository[entities.Room]
}

func NewRoomRepository(collection *mongo.Collection) RoomRepositoryInterface {
	return &roomRepository{
		base: NewRepository[entities.Room](collection),
	}
}

func (r *roomRepository) GetBaseRepo() Repository[entities.Room] {
	return r.base
}
