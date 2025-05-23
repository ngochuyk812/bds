package connectrpc

import (
	"context"
	"net/http"
	"property_service/internal/config"
	amenitydto "property_service/internal/dtos/amenity"
	dtos "property_service/internal/dtos/shared"
	"property_service/internal/infra"
	"property_service/internal/infra/global"
	"property_service/internal/usecases"

	"connectrpc.com/connect"
	"github.com/ngochuyk812/building_block/interceptors"
	propertyv1 "github.com/ngochuyk812/proto-bds/gen/property/v1"
	"github.com/ngochuyk812/proto-bds/gen/property/v1/propertyv1connect"
	"github.com/ngochuyk812/proto-bds/gen/statusmsg/v1"
	utilsv1 "github.com/ngochuyk812/proto-bds/gen/utils/v1"
)

// CreateAmenity implements propertyv1connect.PropertyServiceHandler.
func (p *propertyServerHandler) CreateAmenity(ctx context.Context, req *connect.Request[propertyv1.CreateAmenityRequest]) (res *connect.Response[propertyv1.CreateAmenityResponse], err error) {
	res = connect.NewResponse(&propertyv1.CreateAmenityResponse{
		Status: &statusmsg.StatusMessage{},
	})
	dto := amenitydto.CreateAmenityRequest{
		Name:        req.Msg.GetName(),
		Description: req.Msg.GetDescription(),
		Icon:        req.Msg.GetIcon(),
	}
	if err := global.Validate.Struct(dto); err != nil {
		res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_VALIDATION_FAILED
		res.Msg.Status.Extras = []string{err.Error()}
		return res, nil
	}
	result, err := p.useCases.GetAmenitiesUseCase().CreateAmenity(ctx, dto)

	if err != nil {
		res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_INTERNAL_ERROR
		return res, err
	}

	res.Msg.Status.Code = result.Code
	res.Msg.Status.Extras = result.Extras
	return res, nil

}

// DeleteAmenity implements propertyv1connect.PropertyServiceHandler.
func (p *propertyServerHandler) DeleteAmenity(ctx context.Context, req *connect.Request[propertyv1.DeleteAmenityRequest]) (res *connect.Response[propertyv1.DeleteAmenityResponse], err error) {
	res = connect.NewResponse(&propertyv1.DeleteAmenityResponse{
		Status: &statusmsg.StatusMessage{},
	})
	dto := amenitydto.DeleteAmenityRequest{
		Guid: req.Msg.GetGuid(),
	}
	if err := global.Validate.Struct(dto); err != nil {
		res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_VALIDATION_FAILED
		res.Msg.Status.Extras = []string{err.Error()}
		return res, nil
	}
	result, err := p.useCases.GetAmenitiesUseCase().DeleteAmenity(ctx, dto)

	if err != nil {
		res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_INTERNAL_ERROR
		return res, err
	}
	res.Msg.Status.Code = result.Code
	res.Msg.Status.Extras = result.Extras

	return res, nil
}

// SearchAdvanceAmenities implements propertyv1connect.PropertyServiceHandler.
func (p *propertyServerHandler) SearchAdvanceAmenities(ctx context.Context, req *connect.Request[utilsv1.SearchAdvanceRequest]) (*connect.Response[propertyv1.SearchAdvanceAmenitiesResponse], error) {
	res := connect.NewResponse(&propertyv1.SearchAdvanceAmenitiesResponse{
		Status: &statusmsg.StatusMessage{},
	})
	query := dtos.SearchAdvanceModel{
		Filters: make(map[string]dtos.FilterModel),
	}

	for _, filter := range req.Msg.GetFilters() {
		query.Filters[filter.GetField()] = dtos.FilterModel{
			Type:   filter.GetType(),
			Filter: filter.GetFilter(),
		}
	}
	for _, sort := range req.Msg.GetSort() {
		query.Sort = append(query.Sort, dtos.SortModelItem{
			ColId: sort.GetColId(),
			Sort:  sort.GetSort(),
		})
	}

	result, err := p.cabin.GetUnitOfWork().GetAmenityRepository().GetBaseRepo().SearchAdvance(ctx, query)
	if err != nil {
		res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_VALIDATION_FAILED
		return res, err
	}
	res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	res.Msg.Total = int64(result.Total)
	res.Msg.Rows = make([]*propertyv1.AmenityModel, len(result.Rows))
	for i, item := range result.Rows {
		res.Msg.Rows[i] = &propertyv1.AmenityModel{
			Guid:        item.Guid,
			Name:        item.Name,
			Description: item.Description,
			Icon:        item.Icon,
		}
	}
	return res, nil

}

// UpdateAmenity implements propertyv1connect.PropertyServiceHandler.
func (p *propertyServerHandler) UpdateAmenity(ctx context.Context, req *connect.Request[propertyv1.UpdateAmenityRequest]) (res *connect.Response[propertyv1.UpdateAmenityResponse], err error) {
	res = connect.NewResponse(&propertyv1.UpdateAmenityResponse{
		Status: &statusmsg.StatusMessage{},
	})
	dto := amenitydto.UpdateAmenityRequest{
		Name:        req.Msg.GetName(),
		Description: req.Msg.GetDescription(),
		Icon:        req.Msg.GetIcon(),
		Guid:        req.Msg.GetGuid(),
	}
	if err := global.Validate.Struct(dto); err != nil {
		res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_VALIDATION_FAILED
		res.Msg.Status.Extras = []string{err.Error()}
		return res, nil
	}
	result, err := p.useCases.GetAmenitiesUseCase().UpdateAmenity(ctx, dto)

	if err != nil {
		res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_INTERNAL_ERROR
		return res, err
	}
	res.Msg.Status.Extras = result.Extras
	res.Msg.Status.Code = result.Code
	return res, nil
}

// FetchAmenities implements propertyv1connect.PropertyServiceHandler.
func (p *propertyServerHandler) FetchAmenities(context.Context, *connect.Request[propertyv1.FetchAmenitiesRequest]) (*connect.Response[propertyv1.FetchAmenitiesResponse], error) {
	panic("unimplemented")
}

func NewPropertyServer(cabin infra.Cabin, useCases usecases.UsecaseManager) (pattern string, handler http.Handler) {
	impl := &propertyServerHandler{
		cabin:    cabin,
		useCases: useCases,
	}
	path, handler := propertyv1connect.NewPropertyServiceHandler(impl,
		connect.WithInterceptors(
			interceptors.NewLoggingInterceptor(),
			interceptors.NewAuthInterceptor(config.SecretKey, &global.PoliciesPath),
		),
		connect.WithIdempotency(connect.IdempotencyIdempotent),
	)
	return path, handler
}
