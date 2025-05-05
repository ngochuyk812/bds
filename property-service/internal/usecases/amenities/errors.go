package usercase_amenities

import "errors"

var (
	ErrAmenityExist     = errors.New("amenity already exist")
	ErrAmenityNotFound  = errors.New("amenity not found")
	ErrAmenityNotActive = errors.New("amenity not active")
	ErrInvalidRequest   = errors.New("invalid request")
)
