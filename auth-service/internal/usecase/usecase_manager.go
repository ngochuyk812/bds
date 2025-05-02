package usecase

import (
	"auth_service/internal/infra"
	siteusecase "auth_service/internal/usecase/site"
	userusecase "auth_service/internal/usecase/user"
)

type UsecaseManager interface {
	GetSiteUseCase() siteusecase.SiteUseCase
	GetUserUsecase() userusecase.UserUsecase
}

type usecaseManager struct {
	siteUseCase siteusecase.SiteUseCase
	userUsecase userusecase.UserUsecase
}

func NewUsecaseManager(cabin infra.Cabin) UsecaseManager {
	return &usecaseManager{
		siteUseCase: siteusecase.NewSiteUseCase(cabin),
		userUsecase: userusecase.NewUserUsecase(cabin),
	}
}

func (u *usecaseManager) GetSiteUseCase() siteusecase.SiteUseCase {
	return u.siteUseCase
}

func (u *usecaseManager) GetUserUsecase() userusecase.UserUsecase {
	return u.userUsecase
}
