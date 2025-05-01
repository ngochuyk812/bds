package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Area struct {
	BaseEntity  `bson:",inline"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	CityID      primitive.ObjectID `json:"city_id" bson:"city_id,omitempty"`
}
