package connectrpc

import (
	sitedto "auth_service/internal/dtos/site"
	"context"

	"connectrpc.com/connect"
	authv1 "github.com/ngochuyk812/proto-bds/gen/auth/v1"
	"github.com/ngochuyk812/proto-bds/gen/statusmsg/v1"
	utilsv1 "github.com/ngochuyk812/proto-bds/gen/utils/v1"
)

func (s *authServerHandler) CreateSite(ctx context.Context, req *connect.Request[authv1.CreateSiteRequest]) (res *connect.Response[authv1.CreateSiteResponse], err error) {
	res = connect.NewResponse(&authv1.CreateSiteResponse{
		Status: &statusmsg.StatusMessage{},
	})

	err = s.usecaseManager.GetSiteUseCase().CreateSite(ctx, sitedto.CreateSiteCommand{
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

	err = s.usecaseManager.GetSiteUseCase().DeleteSite(ctx, sitedto.DeleteSiteCommand{
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

	pageSize := int32(req.Msg.Pagination.GetPageSize())
	if pageSize <= 0 {
		pageSize = 10
	}

	pageNumber := int32(req.Msg.Pagination.GetPageNumber())
	if pageNumber <= 0 {
		pageNumber = 1
	}

	result, err := s.usecaseManager.GetSiteUseCase().GetSitesPaging(ctx, sitedto.FetchSitesQuery{
		Page:     pageNumber,
		PageSize: pageSize,
	})

	if err != nil {
		res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}

	items := make([]*authv1.SiteModel, len(result.Items))
	for i, item := range result.Items {
		items[i] = &authv1.SiteModel{
			Id:     item.ID,
			Guid:   item.Guid,
			Name:   item.Name,
			SiteId: item.SiteId,
		}
	}

	res.Msg.Items = items
	res.Msg.Pagination = &utilsv1.PaginationResponse{
		CurrentPage: req.Msg.Pagination.PageNumber,
		PageSize:    req.Msg.Pagination.PageSize,
		Total:       int64(result.Total),
		TotalPages:  int64(result.TotalPages),
	}

	res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	return res, nil
}

func (s *authServerHandler) UpdateSite(ctx context.Context, req *connect.Request[authv1.UpdateSiteRequest]) (res *connect.Response[authv1.UpdateSiteResponse], err error) {
	res = connect.NewResponse(&authv1.UpdateSiteResponse{
		Status: &statusmsg.StatusMessage{},
	})

	err = s.usecaseManager.GetSiteUseCase().UpdateSite(ctx, sitedto.UpdateSiteCommand{
		Name: req.Msg.GetName(),
		Guid: req.Msg.GetGuid(),
	})

	if err != nil {
		res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}

	res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	return res, nil
}
