package usecase

import "auth_service/internal/infra"

type UsecaseManager interface {
	GetSiteUseCase() SiteUseCase
	GetUserUsecase() UserUsecase
}

type usecaseManager struct {
	siteUseCase SiteUseCase
	userUsecase UserUsecase
}

func NewUsecaseManager(cabin infra.Cabin) UsecaseManager {
	return &usecaseManager{
		siteUseCase: NewSiteUseCase(cabin),
		userUsecase: NewUserUsecase(cabin),
	}
}

func (u *usecaseManager) GetSiteUseCase() SiteUseCase {
	return u.siteUseCase
}

func (u *usecaseManager) GetUserUsecase() UserUsecase {
	return u.userUsecase
}
