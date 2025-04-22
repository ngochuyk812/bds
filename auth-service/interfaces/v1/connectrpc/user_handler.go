package connectrpc

import (
	commands_auth "auth_service/internal/app/commands/auth"
	commands_user "auth_service/internal/app/commands/user"
	"context"

	"connectrpc.com/connect"
	bus_core "github.com/ngochuyk812/building_block/pkg/mediator/bus"
	authv1 "github.com/ngochuyk812/proto-bds/gen/auth/v1"
	"github.com/ngochuyk812/proto-bds/gen/statusmsg/v1"
)

// GetProfile implements authv1connect.AuthServiceHandler.
func (s *authServerHandler) GetProfile(ctx context.Context, req *connect.Request[authv1.GetProfileRequest]) (*connect.Response[authv1.GetProfileResponse], error) {
	res := connect.NewResponse(&authv1.GetProfileResponse{
		Status: &statusmsg.StatusMessage{},
	})

	result, err := bus_core.Send[commands_user.GetProfileCommand, commands_user.GetProfileCommandResponse](
		s.cabin.GetInfra().GetMediator(),
		ctx,
		commands_user.GetProfileCommand{},
	)
	if err != nil {
		res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}

	res.Msg.Status = result.StatusMessage
	res.Msg.Profile = &authv1.UserDetail{
		FirstName: result.FirstName,
		LastName:  result.LastName,
		Phone:     result.Phone,
		Address:   result.Address,
	}
	return res, nil
}

// Logout implements authv1connect.AuthServiceHandler.
func (s *authServerHandler) Logout(ctx context.Context, req *connect.Request[authv1.LogoutRequest]) (*connect.Response[authv1.LogoutResponse], error) {
	res := connect.NewResponse(&authv1.LogoutResponse{
		Status: &statusmsg.StatusMessage{},
	})

	result, err := bus_core.Send[commands_auth.LogoutCommand, commands_auth.LogoutCommandResponse](
		s.cabin.GetInfra().GetMediator(),
		ctx,
		commands_auth.LogoutCommand{},
	)
	if err != nil {
		res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}

	res.Msg.Status = result.StatusMessage
	return res, nil
}

func (s *authServerHandler) UpdateProfile(ctx context.Context, req *connect.Request[authv1.UpdateProfileRequest]) (*connect.Response[authv1.UpdateProfileResponse], error) {
	res := connect.NewResponse(&authv1.UpdateProfileResponse{
		Status: &statusmsg.StatusMessage{},
	})

	result, err := bus_core.Send[commands_user.UpdateProfileCommand, commands_user.UpdateProfileCommandResponse](
		s.cabin.GetInfra().GetMediator(),
		ctx,
		commands_user.UpdateProfileCommand{
			FirstName: req.Msg.GetFirstName(),
			LastName:  req.Msg.GetLastName(),
			Phone:     req.Msg.GetPhone(),
			Address:   req.Msg.GetAddress(),
		},
	)
	if err != nil {
		res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}

	res.Msg.Status = result.StatusMessage
	return res, nil
}
