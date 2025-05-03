package amenitydto

import "github.com/ngochuyk812/proto-bds/gen/statusmsg/v1"

type CreateAmenityCommand struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Icon        string `json:"icon" validate:"required"`
}

type CreateAmenityCommandResponse struct {
	*statusmsg.StatusMessage
}

type UpdateAmenityCommand struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"Amenity_id" validate:"required"`
	Icon        string `json:"icon" validate:"required"`
	Guid        string `json:"guid" validate:"required"`
}
type UpdateAmenityCommandResponse struct {
	*statusmsg.StatusMessage
}

type DeleteAmenityCommand struct {
	Guid string `json:"guid" validate:"required"`
}
type DeleteAmenityCommandResponse struct {
	*statusmsg.StatusMessage
}

type FetchAmenitiesQuery struct {
	Page     int32
	PageSize int32
}

type AmenityModel struct {
	Guid        string `json:"guid"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"Amenity_id" validate:"required"`
	Icon        string `json:"icon" validate:"required"`
}

type FetchAmenitiesResponse struct {
	Items      []AmenityModel
	Total      int
	Page       int32
	PageSize   int32
	TotalPages int32
}
