package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Room struct {
	BaseEntity  `bson:",inline"`
	Name        string               `json:"name" bson:"name"`
	Description string               `json:"description" bson:"description"`
	Area        float64              `json:"area" bson:"area"`
	PropertyID  primitive.ObjectID   `json:"property_id" bson:"property_id,omitempty"`
	RoomTypeID  primitive.ObjectID   `json:"room_type_id" bson:"room_type_id,omitempty"`
	MediaIDs    []primitive.ObjectID `json:"media_ids" bson:"media_ids,omitempty"`
	CategoryIDs []primitive.ObjectID `json:"category_ids" bson:"category_ids,omitempty"`
	AmenityIDs  []primitive.ObjectID `json:"amenity_ids" bson:"amenity_ids,omitempty"`
	Status      string               `json:"status" bson:"status"`
}
