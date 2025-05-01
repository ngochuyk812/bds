package entities

type RoomType struct {
	BaseEntity  `bson:",inline"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}
