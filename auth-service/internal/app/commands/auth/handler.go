package commands_auth

import (
	"auth_service/internal/app/eventbus/events"
	"auth_service/internal/entity"
	"auth_service/internal/infra"
	"auth_service/pkg"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/Microsoft/go-winio/pkg/guid"
	"github.com/ngochuyk812/building_block/infrastructure/helpers"
	auth_context "github.com/ngochuyk812/building_block/pkg/auth"
	bus_core "github.com/ngochuyk812/building_block/pkg/mediator/bus"
	"github.com/ngochuyk812/proto-bds/gen/statusmsg/v1"
)

const (
	KEY_CACHE_OTP_SIGNUP    = "KEY_CACHE_OTP_SIGNUP"
	KEY_CACHE_REFRESH_TOKEN = "REFRESH_TOKEN"
)

type LoginHandler struct {
	Cabin infra.Cabin
}

var _ bus_core.IHandler[LoginCommand, LoginCommandResponse] = (*LoginHandler)(nil)

func (h *LoginHandler) Handle(ctx context.Context, cmd LoginCommand) (LoginCommandResponse, error) {
	res := LoginCommandResponse{
		StatusMessage: &statusmsg.StatusMessage{},
	}
	exist, err := h.Cabin.GetUnitOfWork().GetUserRepository().GetUserByEmail(ctx, cmd.Email)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	if exist == nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_USER_NOT_EXIST
		return res, nil
	}
	if exist.Active.Bool == false {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_USER_NOT_ACTIVE
		return res, nil
	}
	verify := pkg.VerifyHashPassword(cmd.Password, exist.HashPassword, exist.Salt)

	if verify == false {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_INCORRECT_PASSWORD
		return res, nil
	}
	token, err := auth_context.GenerateJWT(&auth_context.ClaimModel{
		IdSite:     exist.Siteid,
		IdAuthUser: exist.Guid,
		Roles:      []string{"user"},
		UserName:   exist.Email,
		Email:      exist.Email,
	}, h.Cabin.GetInfra().GetConfig().SecretKey, 1*time.Hour)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	ttlRefreshToken := 30 * 24 * time.Hour
	refreshToken, err := auth_context.GenerateJWT(&auth_context.ClaimModel{
		IdSite:     exist.Siteid,
		IdAuthUser: exist.Guid,
	}, h.Cabin.GetInfra().GetConfig().SecretKey+"_REFRESH_TOKEN", ttlRefreshToken)

	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	cache := h.Cabin.GetInfra().GetCache()
	err = cache.Set(ctx, cache.WithPrefix(KEY_CACHE_REFRESH_TOKEN, exist.Guid), refreshToken, ttlRefreshToken)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	res.AccessToken = token
	res.RefreshToken = refreshToken
	res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	return res, nil
}

type SignUpCommandHandler struct {
	Cabin infra.Cabin
}

var _ bus_core.IHandler[SignUpCommand, SignUpCommandResponse] = (*SignUpCommandHandler)(nil)

func (h *SignUpCommandHandler) Handle(ctx context.Context, cmd SignUpCommand) (SignUpCommandResponse, error) {
	res := SignUpCommandResponse{
		StatusMessage: &statusmsg.StatusMessage{},
	}
	cache := h.Cabin.GetInfra().GetCache()

	exist, err := h.Cabin.GetUnitOfWork().GetUserRepository().GetUserByEmail(ctx, cmd.Email)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	if exist != nil && exist.Active.Bool == true {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_USER_EXIST
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
	if exist != nil {
		userEntity := &entity.User{
			Guid:         exist.Guid,
			Email:        exist.Email,
			HashPassword: hash,
			Salt:         salt,
			Active:       exist.Active,
			Updatedat:    sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
		}
		err = h.Cabin.GetUnitOfWork().GetUserRepository().UpdateUser(ctx, userEntity)
	} else {
		userEntity := &entity.User{
			Guid:         guid.String(),
			Email:        cmd.Email,
			HashPassword: hash,
			Salt:         salt,
			Createdat:    time.Now().Unix(),
		}
		err = h.Cabin.GetUnitOfWork().GetUserRepository().CreateUser(ctx, userEntity)
	}

	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}

	otp := fmt.Sprintf("%06d", rand.Intn(1000000))

	err = cache.Set(ctx, cache.WithPrefix(KEY_CACHE_OTP_SIGNUP, cmd.Email), otp, 10*time.Minute)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}

	h.Cabin.GetInfra().GetEventbus().Publish(ctx, &events.IntegrationEventSendMail{
		Body:    "OTP của bạn là " + otp,
		Subject: "Mã xác nhận đăng ký",
		To:      []string{cmd.Email},
	})
	res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	return res, nil
}

type VerifySignUpCommandHandler struct {
	Cabin infra.Cabin
}

var _ bus_core.IHandler[VerifySignUpCommand, VerifySignUpCommandResponse] = (*VerifySignUpCommandHandler)(nil)

func (h *VerifySignUpCommandHandler) Handle(ctx context.Context, cmd VerifySignUpCommand) (VerifySignUpCommandResponse, error) {
	res := VerifySignUpCommandResponse{
		StatusMessage: &statusmsg.StatusMessage{},
	}
	cache := h.Cabin.GetInfra().GetCache()
	exist, err := h.Cabin.GetUnitOfWork().GetUserRepository().GetUserByEmail(ctx, cmd.Email)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	if exist == nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_USER_NOT_EXIST
		return res, nil
	}
	if exist.Active.Bool {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_USER_IS_ACTIVE
		return res, nil
	}
	var otp string
	err = cache.Get(ctx, cache.WithPrefix(KEY_CACHE_OTP_SIGNUP, cmd.Email), &otp)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, nil
	}
	if otp != cmd.Otp {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_INCORRECT_OTP
		return res, nil
	}
	userEntity := &entity.User{
		Guid:      exist.Guid,
		Email:     exist.Email,
		Active:    sql.NullBool{Bool: true, Valid: true},
		Updatedat: sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
	}
	err = h.Cabin.GetUnitOfWork().GetUserRepository().UpdateUser(ctx, userEntity)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, nil
	}
	res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	return res, nil
}

type RefreshTokenCommandHandler struct {
	Cabin infra.Cabin
}

var _ bus_core.IHandler[RefreshTokenCommand, RefreshTokenCommandResponse] = (*RefreshTokenCommandHandler)(nil)

func (h *RefreshTokenCommandHandler) Handle(ctx context.Context, cmd RefreshTokenCommand) (RefreshTokenCommandResponse, error) {
	res := RefreshTokenCommandResponse{
		StatusMessage: &statusmsg.StatusMessage{},
	}
	cache := h.Cabin.GetInfra().GetCache()

	claims, err := auth_context.VerifyJWT(cmd.RefreshToken, h.Cabin.GetInfra().GetConfig().SecretKey+"_REFRESH_TOKEN")

	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_REFRESH_TOKEN_INVALID
		return res, err
	}
	exist, err := h.Cabin.GetUnitOfWork().GetUserRepository().GetUserByGuid(ctx, claims.IdAuthUser)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	if exist == nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_USER_NOT_EXIST
		return res, nil
	}
	if exist.Active.Bool == false {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_USER_NOT_ACTIVE
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
	ttlRefreshToken := 30 * 24 * time.Hour
	refreshToken, err := auth_context.GenerateJWT(&auth_context.ClaimModel{
		IdSite:     exist.Siteid,
		IdAuthUser: exist.Guid,
	}, h.Cabin.GetInfra().GetConfig().SecretKey+"_REFRESH_TOKEN", ttlRefreshToken)

	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	err = cache.Set(ctx, cache.WithPrefix(KEY_CACHE_REFRESH_TOKEN, exist.Guid), refreshToken, ttlRefreshToken)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	res.RefreshToken = refreshToken
	res.AccessToken = token
	res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	return res, nil
}

type LogoutHandler struct {
	Cabin infra.Cabin
}

var _ bus_core.IHandler[LogoutCommand, LogoutCommandResponse] = (*LogoutHandler)(nil)

func (h *LogoutHandler) Handle(ctx context.Context, cmd LogoutCommand) (LogoutCommandResponse, error) {
	res := LogoutCommandResponse{
		StatusMessage: &statusmsg.StatusMessage{},
	}

	authContext, ok := helpers.AuthContext(ctx)
	if !ok {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, errors.New("cannot get auth context")
	}

	cache := h.Cabin.GetInfra().GetCache()
	err := cache.Del(ctx, cache.WithPrefix(KEY_CACHE_REFRESH_TOKEN, authContext.IdAuthUser))
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}

	res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	return res, nil
}
