package commands_auth

import (
	"auth_service/internal/infra"
	usercase "auth_service/internal/usecase/user"
	"auth_service/pkg"
	"context"
	"time"

	"github.com/Microsoft/go-winio/pkg/guid"
	auth_context "github.com/ngochuyk812/building_block/pkg/auth"
	bus_core "github.com/ngochuyk812/building_block/pkg/mediator/bus"
	"github.com/ngochuyk812/proto-bds/gen/statusmsg/v1"
)

type LoginHandler struct {
	Cabin infra.Cabin
}

var _ bus_core.IHandler[LoginCommand, LoginCommandResponse] = (*LoginHandler)(nil)

func (h *LoginHandler) Handle(ctx context.Context, cmd LoginCommand) (LoginCommandResponse, error) {
	res := LoginCommandResponse{}
	exist, err := h.Cabin.GetUnitOfWork().GetUserRepository().GetUserByEmail(ctx, usercase.GetUserByEmailUsecase{
		Email: cmd.Email,
	})
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	if exist == nil {
		// res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_USER_NOT_EXIST
		return res, nil
	}
	verify := pkg.VerifyHashPassword(cmd.Password, exist.HashPassword, exist.Salt)

	if verify == false {
		// res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_INCORRECT_PASSWORD
		return res, nil
	}
	token, err := auth_context.GenerateJWT(&auth_context.ClaimModel{
		IdSite:     exist.Siteid,
		IdAuthUser: exist.Guid,
		Roles:      []string{"user"},
		UserName:   exist.Email,
		Email:      exist.Email,
	}, h.Cabin.GetInfra().GetConfig().SecretKey, 30*time.Millisecond)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}

	res.AccessToken = token
	return res, nil
}

type SignUpCommandHandler struct {
	Cabin infra.Cabin
}

var _ bus_core.IHandler[SignUpCommand, SignUpCommandResponse] = (*SignUpCommandHandler)(nil)

func (h *SignUpCommandHandler) Handle(ctx context.Context, cmd SignUpCommand) (SignUpCommandResponse, error) {
	res := SignUpCommandResponse{}
	exist, err := h.Cabin.GetUnitOfWork().GetUserRepository().GetUserByEmail(ctx, usercase.GetUserByEmailUsecase{
		Email: cmd.Email,
	})
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	if exist != nil {
		// res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_USER_EXIST
		return res, nil
	}
	hash, salt, err := pkg.GenerateHashPassword(cmd.Password)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	guid, err := guid.NewV4()
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	err = h.Cabin.GetUnitOfWork().GetUserRepository().CreateUser(ctx, &usercase.CreateUserUsercase{
		Guid:         guid.String(),
		Email:        cmd.Email,
		HashPassword: hash,
		Salt:         salt,
	})
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	return res, nil
}
