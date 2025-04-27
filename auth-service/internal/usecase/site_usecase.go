package usecase

import (
	sitedto "auth_service/internal/dtos/site"
	"auth_service/internal/entities"
	"auth_service/internal/infra"
	"auth_service/internal/repository"
	"context"
	"time"
)

type siteUseCase struct {
	Cabin infra.Cabin
}

type SiteUseCase interface {
	CreateSite(ctx context.Context, req sitedto.CreateSiteCommand) error

	UpdateSite(ctx context.Context, req sitedto.UpdateSiteCommand) error

	DeleteSite(ctx context.Context, req sitedto.DeleteSiteCommand) error
}

func NewSiteUseCase(cabin infra.Cabin) SiteUseCase {
	return &siteUseCase{
		Cabin: cabin,
	}
}

func (s *siteUseCase) CreateSite(ctx context.Context, req sitedto.CreateSiteCommand) error {
	return s.Cabin.GetUnitOfWork().ExecTx(ctx, func(uow repository.UnitOfWork) error {
		siteEntity := &entities.Site{
			Name: req.Name,
			BaseEntity: entities.BaseEntity{
				SiteId:    req.SiteId,
				CreatedAt: time.Now().Unix(),
			},
		}
		err := uow.GetSiteRepository().CreateSite(ctx, siteEntity)
		return err
	})
}

func (s *siteUseCase) UpdateSite(ctx context.Context, req sitedto.UpdateSiteCommand) error {

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

func (s *siteUseCase) DeleteSite(ctx context.Context, req sitedto.DeleteSiteCommand) error {
	return s.Cabin.GetUnitOfWork().ExecTx(ctx, func(uow repository.UnitOfWork) error {
		err := uow.GetSiteRepository().DeleteSiteByGuid(ctx, req.Guid)
		return err
	})
}
