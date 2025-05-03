package usecases

import (
	"property_service/internal/infra"
	usercase_amenities "property_service/internal/usecases/amenities"
)

type UsecaseManager interface {
	GetAmenitiesUseCase() usercase_amenities.AmenityUseCase
}

type usecaseManager struct {
	amenitiyUseCases usercase_amenities.AmenityUseCase
}

func NewUsecaseManager(cabin infra.Cabin) UsecaseManager {
	return &usecaseManager{
		amenitiyUseCases: usercase_amenities.NewAmenityUseCase(cabin),
	}
}

func (u *usecaseManager) GetAmenitiesUseCase() usercase_amenities.AmenityUseCase {
	return u.amenitiyUseCases
}
