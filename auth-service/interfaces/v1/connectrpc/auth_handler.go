package connectrpc

import (
	userdto "auth_service/internal/dtos/user"
	"auth_service/internal/infra/global"
	"context"

	"connectrpc.com/connect"
	authv1 "github.com/ngochuyk812/proto-bds/gen/auth/v1"
	"github.com/ngochuyk812/proto-bds/gen/statusmsg/v1"
)

func (s *authServerHandler) Login(ctx context.Context, req *connect.Request[authv1.LoginRequest]) (*connect.Response[authv1.LoginResponse], error) {
	res := connect.NewResponse(&authv1.LoginResponse{
		Status: &statusmsg.StatusMessage{},
	})
	dto := userdto.LoginCommand{
		Email:    req.Msg.GetEmail(),
		Password: req.Msg.GetPassword(),
	}

	if err := global.Validate.Struct(dto); err != nil {
		res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_VALIDATION_FAILED
		res.Msg.Status.Extras = []string{err.Error()}
		return res, nil
	}
	result, err := s.usecaseManager.GetUserUsecase().Login(ctx, dto)

	if err != nil {
		res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}

	res.Msg.AccessToken = result.AccessToken
	res.Msg.RefreshToken = result.RefreshToken
	res.Msg.Status = result.StatusMessage
	return res, nil
}

// SignUp implements authv1connect.AuthServiceHandler.
func (s *authServerHandler) SignUp(ctx context.Context, req *connect.Request[authv1.SignUpRequest]) (*connect.Response[authv1.SignUpResponse], error) {
	res := connect.NewResponse(&authv1.SignUpResponse{
		Status: &statusmsg.StatusMessage{},
	})
	dto := userdto.SignUpCommand{
		Email:    req.Msg.GetEmail(),
		Password: req.Msg.GetPassword(),
	}

	if err := global.Validate.Struct(dto); err != nil {
		res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_VALIDATION_FAILED
		res.Msg.Status.Extras = []string{err.Error()}
		return res, nil
	}

	result, err := s.usecaseManager.GetUserUsecase().SignUp(ctx, dto)

	if err != nil {
		res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}

	res.Msg.Status = result.StatusMessage
	return res, nil
}

// VerifySignUp implements authv1connect.AuthServiceHandler.
func (s *authServerHandler) VerifySignUp(ctx context.Context, req *connect.Request[authv1.VerifySignUpRequest]) (*connect.Response[authv1.VerifySignUpResponse], error) {
	res := connect.NewResponse(&authv1.VerifySignUpResponse{
		Status: &statusmsg.StatusMessage{},
	})
	dto := userdto.VerifySignUpCommand{
		Email: req.Msg.GetEmail(),
		Otp:   req.Msg.GetOtp(),
	}

	if err := global.Validate.Struct(dto); err != nil {
		res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_VALIDATION_FAILED
		res.Msg.Status.Extras = []string{err.Error()}
		return res, nil
	}

	result, err := s.usecaseManager.GetUserUsecase().VerifySignUp(ctx, dto)

	if err != nil {
		res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}

	res.Msg.Status = result.StatusMessage
	return res, nil
}

// RefreshToken implements authv1connect.AuthServiceHandler.
func (s *authServerHandler) RefreshToken(ctx context.Context, req *connect.Request[authv1.RefreshTokenRequest]) (*connect.Response[authv1.RefreshTokenResponse], error) {
	res := connect.NewResponse(&authv1.RefreshTokenResponse{
		Status: &statusmsg.StatusMessage{},
	})

	result, err := s.usecaseManager.GetUserUsecase().RefreshToken(ctx, userdto.RefreshTokenCommand{
		RefreshToken: req.Msg.GetRefreshToken(),
	})

	if err != nil {
		res.Msg.Status.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}

	res.Msg.AccessToken = result.AccessToken
	res.Msg.RefreshToken = result.RefreshToken
	res.Msg.Status = result.StatusMessage
	return res, nil
}
