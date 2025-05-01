package repositories

import (
	"sender_service/internal/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type RoomTypeRepositoryInterface interface {
	GetBaseRepo() *Repository[entities.RoomType]
}

type RoomTypeRepository struct {
	base *Repository[entities.RoomType]
}

func NewRoomTypeRepository(collection *mongo.Collection) RoomTypeRepositoryInterface {
	return &RoomTypeRepository{
		base: &Repository[entities.RoomType]{
			collection: collection,
		},
	}
}

func (r *RoomTypeRepository) GetBaseRepo() *Repository[entities.RoomType] {
	return r.base
}
