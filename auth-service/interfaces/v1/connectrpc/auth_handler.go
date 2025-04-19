package connectrpc

import (
	commands_auth "auth_service/internal/app/commands/auth"
	"context"

	"connectrpc.com/connect"
	bus_core "github.com/ngochuyk812/building_block/pkg/mediator/bus"
	authv1 "github.com/ngochuyk812/proto-bds/gen/auth/v1"
	"github.com/ngochuyk812/proto-bds/gen/statusmsg/v1"
)

func (s *authServerHandler) Login(ctx context.Context, req *connect.Request[authv1.LoginRequest]) (*connect.Response[authv1.LoginResponse], error) {
	res := connect.NewResponse(&authv1.LoginResponse{
		Status: &statusmsg.StatusMessage{},
	})
	result, err := bus_core.Send[commands_auth.LoginCommand, commands_auth.LoginCommandResponse](s.cabin.GetInfra().GetMediator(), ctx, commands_auth.LoginCommand{
		Email:    req.Msg.GetEmail(),
		Password: req.Msg.GetPassword(),
	})
	if err != nil {
		res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	res.Msg.AccessToken = result.AccessToken
	res.Msg.RefreshToken = result.RefreshToken
	return res, err
}

// SignUp implements authv1connect.AuthServiceHandler.
func (s *authServerHandler) SignUp(ctx context.Context, req *connect.Request[authv1.SignUpRequest]) (*connect.Response[authv1.SignUpResponse], error) {
	res := connect.NewResponse(&authv1.SignUpResponse{
		Status: &statusmsg.StatusMessage{},
	})
	result, err := bus_core.Send[commands_auth.SignUpCommand, commands_auth.SignUpCommandResponse](s.cabin.GetInfra().GetMediator(), ctx, commands_auth.SignUpCommand{
		Email:    req.Msg.GetEmail(),
		Password: req.Msg.GetPassword(),
	})
	if err != nil {
		res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	res.Msg.Status = result.StatusMessage
	return res, err
}

// VerifySignUp implements authv1connect.AuthServiceHandler.
func (s *authServerHandler) VerifySignUp(ctx context.Context, req *connect.Request[authv1.VerifySignUpRequest]) (*connect.Response[authv1.VerifySignUpResponse], error) {
	res := connect.NewResponse(&authv1.VerifySignUpResponse{
		Status: &statusmsg.StatusMessage{},
	})
	result, err := bus_core.Send[commands_auth.VerifySignUpCommand, commands_auth.VerifySignUpCommandResponse](s.cabin.GetInfra().GetMediator(), ctx, commands_auth.VerifySignUpCommand{
		Email: req.Msg.GetEmail(),
		Otp:   req.Msg.GetOtp(),
	})
	if err != nil {
		res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	res.Msg.Status = result.StatusMessage
	return res, err
}
