package connectrpc

import (
	commands_site "auth_service/internal/app/commands/site"
	queries_site "auth_service/internal/app/queries/site"
	"auth_service/internal/infra"
	"context"
	"math"
	"net/http"

	"connectrpc.com/connect"
	"github.com/ngochuyk812/building_block/interceptors"
	"github.com/ngochuyk812/building_block/pkg/dtos"
	bus_core "github.com/ngochuyk812/building_block/pkg/mediator/bus"
	sitev1 "github.com/ngochuyk812/proto-bds/gen/site/v1"
	"github.com/ngochuyk812/proto-bds/gen/site/v1/sitev1connect"
	"github.com/ngochuyk812/proto-bds/gen/statusmsg/v1"
	utilsv1 "github.com/ngochuyk812/proto-bds/gen/utils/v1"
)

var _ sitev1connect.SiteServiceHandler = &siteServerHandler{}

type siteServerHandler struct {
	cabin infra.Cabin
}

func (s *siteServerHandler) CreateSite(ctx context.Context, req *connect.Request[sitev1.CreateSiteRequest]) (res *connect.Response[sitev1.CreateSiteResponse], err error) {
	res = connect.NewResponse[sitev1.CreateSiteResponse](&sitev1.CreateSiteResponse{
		Status: &statusmsg.StatusMessage{},
	})
	_, err = bus_core.Send[commands_site.CreateSiteCommand, commands_site.CreateSiteCommandResponse](s.cabin.GetInfra().GetMediator(), ctx, commands_site.CreateSiteCommand{
		SiteId: req.Msg.GetSiteId(),
		Name:   req.Msg.GetName(),
	})
	if err != nil {
		res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	return res, nil
}

func (s *siteServerHandler) DeleteSite(ctx context.Context, req *connect.Request[sitev1.DeleteSiteRequest]) (res *connect.Response[sitev1.DeleteSiteResponse], err error) {
	res = connect.NewResponse[sitev1.DeleteSiteResponse](&sitev1.DeleteSiteResponse{
		Status: &statusmsg.StatusMessage{},
	})
	_, err = bus_core.Send[commands_site.DeleteSiteCommand, commands_site.DeleteSiteCommandResponse](s.cabin.GetInfra().GetMediator(), ctx, commands_site.DeleteSiteCommand{
		Guid: req.Msg.GetGuid(),
	})
	if err != nil {
		res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	return res, nil
}

func (s *siteServerHandler) FetchSites(ctx context.Context, req *connect.Request[sitev1.FetchSitesRequest]) (res *connect.Response[sitev1.FetchSitesResponse], err error) {
	res = connect.NewResponse[sitev1.FetchSitesResponse](&sitev1.FetchSitesResponse{
		Status: &statusmsg.StatusMessage{},
	})

	paging, err := bus_core.Send[queries_site.FetchSitesQuery, dtos.PagingModel[sitev1.SiteModel]](s.cabin.GetInfra().GetMediator(), ctx, queries_site.FetchSitesQuery{
		PagingRequest: &dtos.PagingRequest{
			PageSize: int(req.Msg.Pagination.GetPageSize()),
			Page:     int(req.Msg.Pagination.GetPageNumber()),
		},
	})
	res.Msg.Pagination = &utilsv1.PaginationResponse{
		CurrentPage: req.Msg.Pagination.PageNumber,
		PageSize:    req.Msg.Pagination.PageSize,
		Total:       int64(paging.Total),
		TotalPages:  int64(math.Ceil(float64(paging.Total) / float64(req.Msg.Pagination.PageSize))),
	}
	items := make([]*sitev1.SiteModel, len(paging.Items))
	for i := range paging.Items {
		items[i] = &paging.Items[i]
	}
	res.Msg.Items = items

	res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	return res, nil
}

func (s *siteServerHandler) UpdateSite(ctx context.Context, req *connect.Request[sitev1.UpdateSiteRequest]) (res *connect.Response[sitev1.UpdateSiteResponse], err error) {
	res = connect.NewResponse[sitev1.UpdateSiteResponse](&sitev1.UpdateSiteResponse{
		Status: &statusmsg.StatusMessage{},
	})
	_, err = bus_core.Send[commands_site.UpdateSiteCommand, commands_site.UpdateSiteCommandResponse](s.cabin.GetInfra().GetMediator(), ctx, commands_site.UpdateSiteCommand{
		SiteId: req.Msg.GetSiteId(),
		Name:   req.Msg.GetName(),
		Guid:   req.Msg.GetGuid(),
	})
	if err != nil {
		res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	return res, nil
}

func NewSiteServer(cabin infra.Cabin) (pattern string, handler http.Handler) {
	impl := &siteServerHandler{
		cabin: cabin,
	}
	path, handler := sitev1connect.NewSiteServiceHandler(impl,
		connect.WithInterceptors(
			interceptors.NewAuthInterceptor(cabin.GetInfra().GetConfig().SecretKey, cabin.GetInfra().GetConfig().PoliciesPath),
			interceptors.NewLoggingInterceptor(cabin.GetInfra().GetLogger()),
		),
		connect.WithIdempotency(connect.IdempotencyIdempotent),
	)
	return path, handler
}
