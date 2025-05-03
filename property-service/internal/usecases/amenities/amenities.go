package usercase_amenities

import "property_service/internal/infra"

type AmenityUseCase interface {
	CrudAmenityUseCase
}
type amenityUseCase struct {
	Cabin infra.Cabin
	CrudAmenityUseCase
}

func NewAmenityUseCase(cabin infra.Cabin) AmenityUseCase {
	return &amenityUseCase{
		Cabin:              cabin,
		CrudAmenityUseCase: NewCrudsAmenityUseCase(cabin),
	}
}
