package queries_site

import (
	"auth_service/internal/infra"
	"context"

	"github.com/ngochuyk812/building_block/pkg/dtos"
	bus_core "github.com/ngochuyk812/building_block/pkg/mediator/bus"
	authv1 "github.com/ngochuyk812/proto-bds/gen/auth/v1"
)

type FetchSitesHandler struct {
	Cabin infra.Cabin
}

var _ bus_core.IHandler[FetchSitesQuery, dtos.PagingModel[authv1.SiteModel]] = (*FetchSitesHandler)(nil)

func (h *FetchSitesHandler) Handle(ctx context.Context, cmd FetchSitesQuery) (dtos.PagingModel[authv1.SiteModel], error) {
	res := dtos.PagingModel[authv1.SiteModel]{}

	paging, err := h.Cabin.GetUnitOfWork().GetSiteRepository().GetSitesPaging(ctx, int32(cmd.PagingRequest.Page), int32(cmd.PagingRequest.PageSize))
	if err != nil {
		return res, err
	}

	res = dtos.PagingModel[authv1.SiteModel]{
		PageSize: cmd.PagingRequest.PageSize,
		Total:    paging.Total,
	}
	res.Items = []authv1.SiteModel{}
	for _, v := range paging.Items {
		res.Items = append(res.Items, authv1.SiteModel{
			Id:     int64(v.ID),
			Guid:   v.Guid,
			Name:   v.Name,
			SiteId: v.Siteid,
		})
	}

	return res, err
}
