package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoomPricing struct {
	BaseEntity  `bson:",inline"`
	RoomID      primitive.ObjectID `json:"room_id" bson:"room_id,omitempty"`
	PricingType string             `json:"pricing_type" bson:"pricing_type"` // "monthly", "daily", etc.
	BasePrice   float64            `json:"base_price" bson:"base_price"`
	Deposit     float64            `json:"deposit" bson:"deposit"`
	Fees        map[string]float64 `json:"fees" bson:"fees"`
	Currency    string             `json:"currency" bson:"currency"`
}
