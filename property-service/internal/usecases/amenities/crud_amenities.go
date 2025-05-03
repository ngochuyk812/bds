package usercase_amenities

import (
	"context"
	"math"
	amenitydto "property_service/internal/dtos/amenity"
	"property_service/internal/entities"
	"property_service/internal/infra"

	"github.com/ngochuyk812/proto-bds/gen/statusmsg/v1"
)

type crudAmenityUseCase struct {
	Cabin infra.Cabin
}

type CrudAmenityUseCase interface {
	CreateAmenity(ctx context.Context, req amenitydto.CreateAmenityRequest) (*amenitydto.CreateAmenityResponse, error)
	UpdateAmenity(ctx context.Context, req amenitydto.UpdateAmenityRequest) (*amenitydto.UpdateAmenityResponse, error)
	DeleteAmenity(ctx context.Context, req amenitydto.DeleteAmenityRequest) (*amenitydto.DeleteAmenityResponse, error)
	GetAmenitiesPaging(ctx context.Context, req amenitydto.FetchAmenitiesRequest) (*amenitydto.FetchAmenitiesResponse, error)
}

func NewCrudsAmenityUseCase(cabin infra.Cabin) CrudAmenityUseCase {
	return &crudAmenityUseCase{
		Cabin: cabin,
	}
}

func (s *crudAmenityUseCase) CreateAmenity(ctx context.Context, req amenitydto.CreateAmenityRequest) (*amenitydto.CreateAmenityResponse, error) {
	res := &amenitydto.CreateAmenityResponse{
		StatusMessage: &statusmsg.StatusMessage{},
	}

	exist, err := s.Cabin.GetUnitOfWork().GetAmenityRepository().GetByName(ctx, req.Name)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_INTERNAL_ERROR
		return res, err
	}
	if exist != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_VALIDATION_FAILED
		res.StatusMessage.Extras = []string{ErrAmenityExist.Error()}
		return res, nil
	}

	amenityEntity := &entities.Amenity{
		Name:        req.Name,
		Description: req.Description,
		Icon:        req.Icon,
		BaseEntity:  entities.NewBaseEntity(ctx),
	}
	amenityEntity, err = s.Cabin.GetUnitOfWork().GetAmenityRepository().GetBaseRepo().Create(ctx, amenityEntity)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	return res, nil
}

func (s *crudAmenityUseCase) UpdateAmenity(ctx context.Context, req amenitydto.UpdateAmenityRequest) (*amenitydto.UpdateAmenityResponse, error) {
	res := &amenitydto.UpdateAmenityResponse{
		StatusMessage: &statusmsg.StatusMessage{},
	}

	exist, err := s.Cabin.GetUnitOfWork().GetAmenityRepository().GetBaseRepo().GetByGuid(ctx, req.Guid)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	if exist == nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_VALIDATION_FAILED
		res.StatusMessage.Extras = []string{ErrAmenityNotFound.Error()}
		return res, nil
	}
	exist.Name = req.Name
	_, err = s.Cabin.GetUnitOfWork().GetAmenityRepository().GetBaseRepo().UpdateByGuid(ctx, req.Guid, exist)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_INTERNAL_ERROR
		return res, err
	}
	res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	return res, nil

}

func (s *crudAmenityUseCase) DeleteAmenity(ctx context.Context, req amenitydto.DeleteAmenityRequest) (*amenitydto.DeleteAmenityResponse, error) {
	res := &amenitydto.DeleteAmenityResponse{
		StatusMessage: &statusmsg.StatusMessage{},
	}
	err := s.Cabin.GetUnitOfWork().GetAmenityRepository().GetBaseRepo().DeleteByGuid(ctx, req.Guid)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	return res, nil
}

func (s *crudAmenityUseCase) GetAmenitiesPaging(ctx context.Context, req amenitydto.FetchAmenitiesRequest) (*amenitydto.FetchAmenitiesResponse, error) {
	paging, err := s.Cabin.GetUnitOfWork().GetAmenityRepository().GetAmenitiesPaging(ctx, "", req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	res := &amenitydto.FetchAmenitiesResponse{
		Page:       req.Page,
		PageSize:   req.PageSize,
		Total:      paging.Total,
		TotalPages: int32(math.Ceil(float64(paging.Total) / float64(req.PageSize))),
	}

	items := make([]amenitydto.AmenityModel, len(paging.Items))
	for i, item := range paging.Items {
		items[i] = amenitydto.AmenityModel{
			Guid:        item.Guid,
			Name:        item.Name,
			Description: item.Description,
			Icon:        item.Icon,
		}
	}
	res.Items = items

	return res, nil
}
