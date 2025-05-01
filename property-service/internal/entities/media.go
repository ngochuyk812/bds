package entities

// Media represents an image, video, or other media file
type Media struct {
	BaseEntity  `bson:",inline"`
	Data        string `json:"data" bson:"data"`
	Description string `json:"description" bson:"description"`
	IsPrimary   bool   `json:"is_primary" bson:"is_primary"`
	Type        string `json:"type" bson:"type"`
}
