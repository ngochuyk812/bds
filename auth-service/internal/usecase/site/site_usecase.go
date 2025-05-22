package siteusecase

import (
	sitedto "auth_service/internal/dtos/site"
	"auth_service/internal/entities"
	"auth_service/internal/infra"
	"context"
	"math"
	"time"

	"github.com/ngochuyk812/proto-bds/gen/statusmsg/v1"
)

type siteUseCase struct {
	Cabin infra.Cabin
}

type SiteUseCase interface {
	CreateSite(ctx context.Context, req sitedto.CreateSiteCommand) (*sitedto.CreateSiteCommandResponse, error)
	UpdateSite(ctx context.Context, req sitedto.UpdateSiteCommand) (*sitedto.UpdateSiteCommandResponse, error)
	DeleteSite(ctx context.Context, req sitedto.DeleteSiteCommand) (*sitedto.DeleteSiteCommandResponse, error)
	GetSitesPaging(ctx context.Context, req sitedto.FetchSitesQuery) (*sitedto.FetchSitesResponse, error)
}

func NewSiteUseCase(cabin infra.Cabin) SiteUseCase {
	return &siteUseCase{
		Cabin: cabin,
	}
}

func (s *siteUseCase) CreateSite(ctx context.Context, req sitedto.CreateSiteCommand) (*sitedto.CreateSiteCommandResponse, error) {
	res := &sitedto.CreateSiteCommandResponse{
		StatusMessage: &statusmsg.StatusMessage{},
	}

	exist, err := s.Cabin.GetUnitOfWork().GetSiteRepository().GetSiteBySiteId(ctx, req.SiteId)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	if exist != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_VALIDATION_FAILED
		res.StatusMessage.Extras = []string{ErrSiteExist.Error()}
		return res, nil
	}

	siteEntity := &entities.Site{
		Name: req.Name,
		BaseEntity: entities.BaseEntity{
			SiteId:    req.SiteId,
			CreatedAt: time.Now().Unix(),
		},
	}
	err = s.Cabin.GetUnitOfWork().GetSiteRepository().GetBaseRepository().Create(ctx, siteEntity)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	return res, nil
}

func (s *siteUseCase) UpdateSite(ctx context.Context, req sitedto.UpdateSiteCommand) (*sitedto.UpdateSiteCommandResponse, error) {
	res := &sitedto.UpdateSiteCommandResponse{
		StatusMessage: &statusmsg.StatusMessage{},
	}

	exist, err := s.Cabin.GetUnitOfWork().GetSiteRepository().GetBaseRepository().GetByGuid(ctx, req.Guid)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	if exist == nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_VALIDATION_FAILED
		res.StatusMessage.Extras = []string{ErrSiteNotFound.Error()}
		return res, nil
	}
	exist.Name = req.Name
	err = s.Cabin.GetUnitOfWork().GetSiteRepository().GetBaseRepository().Update(ctx, exist)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	return res, nil

}

func (s *siteUseCase) DeleteSite(ctx context.Context, req sitedto.DeleteSiteCommand) (*sitedto.DeleteSiteCommandResponse, error) {
	res := &sitedto.DeleteSiteCommandResponse{
		StatusMessage: &statusmsg.StatusMessage{},
	}
	err := s.Cabin.GetUnitOfWork().GetSiteRepository().GetBaseRepository().DeleteByGuid(ctx, req.Guid)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	return res, nil
}

func (s *siteUseCase) GetSitesPaging(ctx context.Context, req sitedto.FetchSitesQuery) (*sitedto.FetchSitesResponse, error) {
	paging, err := s.Cabin.GetUnitOfWork().GetSiteRepository().GetSitesPaging(ctx, req.Page, req.PageSize, req.Name, req.SiteId)
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
