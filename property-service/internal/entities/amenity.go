package entities

type Amenity struct {
	BaseEntity  `bson:",inline"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Icon        string `json:"icon" bson:"icon"`
}
