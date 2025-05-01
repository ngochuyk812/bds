package repositories

import (
	"sender_service/internal/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type RoomRepositoryInterface interface {
	GetBaseRepo() *Repository[entities.Room]
}

type RoomRepository struct {
	base *Repository[entities.Room]
}

func NewRoomRepository(collection *mongo.Collection) RoomRepositoryInterface {
	return &RoomRepository{
		base: &Repository[entities.Room]{
			collection: collection,
		},
	}
}

func (r *RoomRepository) GetBaseRepo() *Repository[entities.Room] {
	return r.base
}
