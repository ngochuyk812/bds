package amenitydto

import "github.com/ngochuyk812/proto-bds/gen/statusmsg/v1"

type CreateAmenityRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Icon        string `json:"icon" validate:"required"`
}

type CreateAmenityResponse struct {
	*statusmsg.StatusMessage
}

type UpdateAmenityRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"Amenity_id" validate:"required"`
	Icon        string `json:"icon" validate:"required"`
	Guid        string `json:"guid" validate:"required"`
}
type UpdateAmenityResponse struct {
	*statusmsg.StatusMessage
}

type DeleteAmenityRequest struct {
	Guid string `json:"guid" validate:"required"`
}
type DeleteAmenityResponse struct {
	*statusmsg.StatusMessage
}

type FetchAmenitiesRequest struct {
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
