package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	BaseEntity  `bson:",inline"`
	Name        string             `json:"name" bson:"name"`
	Icon        string             `json:"icon" bson:"icon"`
	Description string             `json:"description" bson:"description"`
	ParentID    primitive.ObjectID `json:"parent_id" bson:"parent_id,omitempty"`
}
