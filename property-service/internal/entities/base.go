package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseEntity struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	SiteID    string             `json:"site_id" bson:"site_id"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	DeletedAt *time.Time         `json:"deleted_at" bson:"deleted_at"`
}
