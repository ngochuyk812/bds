package usecase

import (
	sitedto "auth_service/internal/dtos/site"
	"auth_service/internal/entities"
	"auth_service/internal/infra"
	"auth_service/internal/repository"
	"context"
	"time"
)

type SiteService struct {
	Cabin infra.Cabin
}

func NewSiteService(cabin infra.Cabin) *SiteService {
	return &SiteService{
		Cabin: cabin,
	}
}

func (s *SiteService) CreateSite(ctx context.Context, name string, siteId string) error {
	return s.Cabin.GetUnitOfWork().ExecTx(ctx, func(uow repository.UnitOfWork) error {
		siteEntity := &entities.Site{
			Name: name,
			BaseEntity: entities.BaseEntity{
				SiteId:    siteId,
				CreatedAt: time.Now().Unix(),
			},
		}
		err := uow.GetSiteRepository().CreateSite(ctx, siteEntity)
		return err
	})
}

func (s *SiteService) UpdateSite(ctx context.Context, req sitedto.UpdateSiteCommand) error {

	exist, err := s.Cabin.GetUnitOfWork().GetSiteRepository().GetSiteByGuid(ctx, req.Guid)
	if exist != nil {
		return nil
	}
	if err != nil {
		return err
	}
	exist.Name = req.Name

	return s.Cabin.GetUnitOfWork().ExecTx(ctx, func(uow repository.UnitOfWork) error {
		err := uow.GetSiteRepository().UpdateSite(ctx, exist)
		return err
	})
}

func (s *SiteService) DeleteSite(ctx context.Context, guid string) error {
	return s.Cabin.GetUnitOfWork().ExecTx(ctx, func(uow repository.UnitOfWork) error {
		err := uow.GetSiteRepository().DeleteSiteByGuid(ctx, guid)
		return err
	})
}
