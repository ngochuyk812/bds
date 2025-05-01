package entities

type PropertyType struct {
	BaseEntity  `bson:",inline"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}
