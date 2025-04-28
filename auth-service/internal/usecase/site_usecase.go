package usecase

import (
	sitedto "auth_service/internal/dtos/site"
	"auth_service/internal/entities"
	"auth_service/internal/infra"
	"auth_service/internal/repository"
	"context"
	"math"
	"time"
)

type siteUseCase struct {
	Cabin infra.Cabin
}

type SiteUseCase interface {
	CreateSite(ctx context.Context, req sitedto.CreateSiteCommand) error
	UpdateSite(ctx context.Context, req sitedto.UpdateSiteCommand) error
	DeleteSite(ctx context.Context, req sitedto.DeleteSiteCommand) error
	GetSitesPaging(ctx context.Context, req sitedto.FetchSitesQuery) (*sitedto.FetchSitesResponse, error)
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
		err := uow.GetSiteRepository().GetBaseRepository().Create(ctx, siteEntity)
		return err
	})
}

func (s *siteUseCase) UpdateSite(ctx context.Context, req sitedto.UpdateSiteCommand) error {

	exist, err := s.Cabin.GetUnitOfWork().GetSiteRepository().GetBaseRepository().GetByGuid(ctx, req.Guid)
	if exist != nil {
		return nil
	}
	if err != nil {
		return err
	}
	exist.Name = req.Name

	return s.Cabin.GetUnitOfWork().ExecTx(ctx, func(uow repository.UnitOfWork) error {
		err := uow.GetSiteRepository().GetBaseRepository().Update(ctx, exist)
		return err
	})
}

func (s *siteUseCase) DeleteSite(ctx context.Context, req sitedto.DeleteSiteCommand) error {
	return s.Cabin.GetUnitOfWork().ExecTx(ctx, func(uow repository.UnitOfWork) error {
		err := uow.GetSiteRepository().GetBaseRepository().DeleteByGuid(ctx, req.Guid)
		return err
	})
}

func (s *siteUseCase) GetSitesPaging(ctx context.Context, req sitedto.FetchSitesQuery) (*sitedto.FetchSitesResponse, error) {
	paging, err := s.Cabin.GetUnitOfWork().GetSiteRepository().GetSitesPaging(ctx, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	res := &sitedto.FetchSitesResponse{
		Page:       req.Page,
		PageSize:   req.PageSize,
		Total:      paging.Total,
		TotalPages: int32(math.Ceil(float64(paging.Total) / float64(req.PageSize))),
	}

	items := make([]sitedto.SiteModel, len(paging.Items))
	for i, item := range paging.Items {
		items[i] = sitedto.SiteModel{
			ID:     int64(item.ID),
			Guid:   item.Guid,
			Name:   item.Name,
			SiteId: item.SiteId,
		}
	}
	res.Items = items

	return res, nil
}
