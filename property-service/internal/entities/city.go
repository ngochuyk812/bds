package entities

type City struct {
	BaseEntity  `bson:",inline"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	CountryCode string `json:"country_code" bson:"country_code"`
}
