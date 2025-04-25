package commands_user

import (
	"auth_service/internal/entity"
	"auth_service/internal/infra"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/ngochuyk812/building_block/infrastructure/helpers"
	bus_core "github.com/ngochuyk812/building_block/pkg/mediator/bus"
	"github.com/ngochuyk812/proto-bds/gen/statusmsg/v1"
)

type UpdateProfileHandler struct {
	Cabin infra.Cabin
}

var _ bus_core.IHandler[UpdateProfileCommand, UpdateProfileCommandResponse] = (*UpdateProfileHandler)(nil)

func (h *UpdateProfileHandler) Handle(ctx context.Context, cmd UpdateProfileCommand) (UpdateProfileCommandResponse, error) {
	res := UpdateProfileCommandResponse{
		StatusMessage: &statusmsg.StatusMessage{},
	}
	authContext, oke := helpers.AuthContext(ctx)
	if !oke {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, errors.New("cannot get auth context")
	}
	exist, err := h.Cabin.GetUnitOfWork().GetUserRepository().GetUserByGuid(ctx, authContext.IdAuthUser)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	if exist == nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_USER_NOT_EXIST
		return res, nil
	}

	userDetailEntity := &entity.UserDetail{
		UserGuid:  authContext.IdAuthUser,
		FirstName: cmd.FirstName,
		LastName:  cmd.LastName,
		Phone:     cmd.Phone,
		Address:   cmd.Address,
		Updatedat: sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
	}
	err = h.Cabin.GetUnitOfWork().GetUserDetailRepository().UpdateUserDetail(ctx, userDetailEntity)

	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}

	res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	return res, nil
}

type GetProfileHandler struct {
	Cabin infra.Cabin
}

var _ bus_core.IHandler[GetProfileCommand, GetProfileCommandResponse] = (*GetProfileHandler)(nil)

func (h *GetProfileHandler) Handle(ctx context.Context, cmd GetProfileCommand) (GetProfileCommandResponse, error) {
	res := GetProfileCommandResponse{
		StatusMessage: &statusmsg.StatusMessage{},
	}

	authContext, ok := helpers.AuthContext(ctx)
	if !ok {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, errors.New("cannot get auth context")
	}

	user, err := h.Cabin.GetUnitOfWork().GetUserRepository().GetUserByGuid(ctx, authContext.IdAuthUser)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	if user == nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_USER_NOT_EXIST
		return res, nil
	}

	userDetail, err := h.Cabin.GetUnitOfWork().GetUserDetailRepository().GetUserDetailByUserGuid(ctx, authContext.IdAuthUser)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}

	res.Email = user.Email
	if userDetail != nil {
		res.FirstName = userDetail.FirstName
		res.LastName = userDetail.LastName
		res.Phone = userDetail.Phone
		res.Address = userDetail.Address
	}

	res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	return res, nil
}
