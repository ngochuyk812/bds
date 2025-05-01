package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Property struct {
	BaseEntity     `bson:",inline"`
	Name           string               `json:"name" bson:"name"`
	Address        string               `json:"address" bson:"address"`
	Description    string               `json:"description" bson:"description"`
	Contact        Contact              `json:"contact" bson:"contact"`
	CityID         primitive.ObjectID   `json:"city_id" bson:"city_id,omitempty"`
	AreaID         primitive.ObjectID   `json:"area_id" bson:"area_id,omitempty"`
	Latitude       float64              `json:"latitude" bson:"latitude"`
	Longitude      float64              `json:"longitude" bson:"longitude"`
	MediaIDs       []primitive.ObjectID `json:"media_ids" bson:"media_ids,omitempty"`
	PropertyTypeID primitive.ObjectID   `json:"property_type_id" bson:"property_type_id,omitempty"`
	IsActive       bool                 `json:"is_active" bson:"is_active"`
}

type Contact struct {
	FullName string `json:"full_name" bson:"full_name"`
	Phone    string `json:"phone" bson:"phone"`
	Email    string `json:"email" bson:"email"`
}
