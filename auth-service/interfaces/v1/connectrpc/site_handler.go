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
	authv1 "github.com/ngochuyk812/proto-bds/gen/auth/v1"
	"github.com/ngochuyk812/proto-bds/gen/auth/v1/authv1connect"
	"github.com/ngochuyk812/proto-bds/gen/statusmsg/v1"
	utilsv1 "github.com/ngochuyk812/proto-bds/gen/utils/v1"
)

func (s *authServerHandler) CreateSite(ctx context.Context, req *connect.Request[authv1.CreateSiteRequest]) (res *connect.Response[authv1.CreateSiteResponse], err error) {
	res = connect.NewResponse(&authv1.CreateSiteResponse{
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

func (s *authServerHandler) DeleteSite(ctx context.Context, req *connect.Request[authv1.DeleteSiteRequest]) (res *connect.Response[authv1.DeleteSiteResponse], err error) {
	res = connect.NewResponse(&authv1.DeleteSiteResponse{
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

func (s *authServerHandler) FetchSites(ctx context.Context, req *connect.Request[authv1.FetchSitesRequest]) (res *connect.Response[authv1.FetchSitesResponse], err error) {
	res = connect.NewResponse(&authv1.FetchSitesResponse{
		Status: &statusmsg.StatusMessage{},
	})

	paging, err := bus_core.Send[queries_site.FetchSitesQuery, dtos.PagingModel[authv1.SiteModel]](s.cabin.GetInfra().GetMediator(), ctx, queries_site.FetchSitesQuery{
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
	items := make([]*authv1.SiteModel, len(paging.Items))
	for i := range paging.Items {
		items[i] = &paging.Items[i]
	}
	res.Msg.Items = items

	res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	return res, nil
}

func (s *authServerHandler) UpdateSite(ctx context.Context, req *connect.Request[authv1.UpdateSiteRequest]) (res *connect.Response[authv1.UpdateSiteResponse], err error) {
	res = connect.NewResponse(&authv1.UpdateSiteResponse{
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

func NewAuthServer(cabin infra.Cabin) (pattern string, handler http.Handler) {
	impl := &authServerHandler{
		cabin: cabin,
	}
	path, handler := authv1connect.NewAuthServiceHandler(impl,
		connect.WithInterceptors(
			interceptors.NewAuthInterceptor(cabin.GetInfra().GetConfig().SecretKey, cabin.GetInfra().GetConfig().PoliciesPath),
			interceptors.NewLoggingInterceptor(cabin.GetInfra().GetLogger()),
		),
		connect.WithIdempotency(connect.IdempotencyIdempotent),
	)
	return path, handler
}
