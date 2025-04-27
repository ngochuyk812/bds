package usecase

import (
	userdto "auth_service/internal/dtos/user"
	"auth_service/internal/entities"
	"auth_service/internal/infra"
	"auth_service/internal/repository"
	"auth_service/pkg"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/Microsoft/go-winio/pkg/guid"
	"github.com/ngochuyk812/building_block/infrastructure/helpers"
	auth_context "github.com/ngochuyk812/building_block/pkg/auth"
	"github.com/ngochuyk812/proto-bds/gen/statusmsg/v1"
)

const (
	KEY_CACHE_OTP_SIGNUP    = "KEY_CACHE_OTP_SIGNUP"
	KEY_CACHE_REFRESH_TOKEN = "REFRESH_TOKEN"
)

type userService struct {
	Cabin infra.Cabin
}

type UserUsecase interface {
	UpdateProfile(ctx context.Context, req userdto.UpdateProfileCommand) (*userdto.UpdateProfileResponse, error)
	GetProfile(ctx context.Context, req userdto.GetProfileCommand) (*userdto.GetProfileResponse, error)
	Login(ctx context.Context, req userdto.LoginCommand) (*userdto.LoginResponse, error)
	SignUp(ctx context.Context, req userdto.SignUpCommand) (*userdto.SignUpResponse, error)
	VerifySignUp(ctx context.Context, req userdto.VerifySignUpCommand) (*userdto.VerifySignUpResponse, error)
	RefreshToken(ctx context.Context, req userdto.RefreshTokenCommand) (*userdto.RefreshTokenResponse, error)
	Logout(ctx context.Context, req userdto.LogoutCommand) (*userdto.LogoutResponse, error)
}

func NewUserUsecase(cabin infra.Cabin) UserUsecase {
	return &userService{
		Cabin: cabin,
	}
}

func (s *userService) UpdateProfile(ctx context.Context, req userdto.UpdateProfileCommand) (*userdto.UpdateProfileResponse, error) {
	res := &userdto.UpdateProfileResponse{
		StatusMessage: &statusmsg.StatusMessage{},
	}

	authContext, ok := helpers.AuthContext(ctx)
	if !ok {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, errors.New("cannot get auth context")
	}

	exist, err := s.Cabin.GetUnitOfWork().GetUserRepository().GetUserByGuid(ctx, authContext.IdAuthUser)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	if exist == nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_USER_NOT_EXIST
		return res, nil
	}

	err = s.Cabin.GetUnitOfWork().ExecTx(ctx, func(uow repository.UnitOfWork) error {
		userDetailEntity := &entities.UserDetail{
			UserGuid:  authContext.IdAuthUser,
			FirstName: &req.FirstName,
			LastName:  &req.LastName,
			Phone:     &req.Phone,
			Address:   &req.Address,
		}
		return uow.GetUserDetailRepository().UpdateUserDetail(ctx, userDetailEntity)
	})

	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}

	res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	return res, nil
}

func (s *userService) GetProfile(ctx context.Context, req userdto.GetProfileCommand) (*userdto.GetProfileResponse, error) {
	res := &userdto.GetProfileResponse{
		StatusMessage: &statusmsg.StatusMessage{},
	}

	authContext, ok := helpers.AuthContext(ctx)
	if !ok {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, errors.New("cannot get auth context")
	}

	user, err := s.Cabin.GetUnitOfWork().GetUserRepository().GetUserByGuid(ctx, authContext.IdAuthUser)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	if user == nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_USER_NOT_EXIST
		return res, nil
	}

	userDetail, err := s.Cabin.GetUnitOfWork().GetUserDetailRepository().GetUserDetailByUserGuid(ctx, authContext.IdAuthUser)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}

	res.Email = user.Email
	if userDetail != nil {
		res.FirstName = *userDetail.FirstName
		res.LastName = *userDetail.LastName
		res.Phone = *userDetail.Phone
		res.Address = *userDetail.Address
	}

	res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	return res, nil
}

func (s *userService) Login(ctx context.Context, req userdto.LoginCommand) (*userdto.LoginResponse, error) {
	res := &userdto.LoginResponse{
		StatusMessage: &statusmsg.StatusMessage{},
	}

	exist, err := s.Cabin.GetUnitOfWork().GetUserRepository().GetUserByEmail(ctx, req.Email)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	if exist == nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_USER_NOT_EXIST
		return res, nil
	}
	if !exist.Active {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_USER_NOT_ACTIVE
		return res, nil
	}

	verify := pkg.VerifyHashPassword(req.Password, exist.HashPassword, exist.Salt)
	if !verify {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_INCORRECT_PASSWORD
		return res, nil
	}

	token, err := auth_context.GenerateJWT(&auth_context.ClaimModel{
		IdSite:     exist.SiteId,
		IdAuthUser: exist.Guid,
		Roles:      []string{"user"},
		UserName:   exist.Email,
		Email:      exist.Email,
	}, s.Cabin.GetInfra().GetConfig().SecretKey, 1*time.Hour)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}

	ttlRefreshToken := 30 * 24 * time.Hour
	refreshToken, err := auth_context.GenerateJWT(&auth_context.ClaimModel{
		IdSite:     exist.SiteId,
		IdAuthUser: exist.Guid,
	}, s.Cabin.GetInfra().GetConfig().SecretKey+"_REFRESH_TOKEN", ttlRefreshToken)

	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}

	cache := s.Cabin.GetInfra().GetCache()
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

func (s *userService) SignUp(ctx context.Context, req userdto.SignUpCommand) (*userdto.SignUpResponse, error) {
	res := &userdto.SignUpResponse{
		StatusMessage: &statusmsg.StatusMessage{},
	}

	cache := s.Cabin.GetInfra().GetCache()

	exist, err := s.Cabin.GetUnitOfWork().GetUserRepository().GetUserByEmail(ctx, req.Email)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	if exist != nil && exist.Active {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_USER_EXIST
		return res, nil
	}

	hash, salt, err := pkg.GenerateHashPassword(req.Password)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}

	err = s.Cabin.GetUnitOfWork().ExecTx(ctx, func(uow repository.UnitOfWork) error {

		newGuid, err := guid.NewV4()
		if err != nil {
			return err
		}

		userEntity := &entities.User{
			Email:        req.Email,
			HashPassword: hash,
			Salt:         salt,
			BaseEntity: entities.BaseEntity{
				Guid: newGuid.String(),
			},
		}
		err = uow.GetUserRepository().CreateUser(ctx, userEntity)

		if err != nil {
			return err
		}

		if req.FirstName != "" || req.LastName != "" {
			userDetailEntity := &entities.UserDetail{
				UserGuid:  newGuid.String(),
				FirstName: &req.FirstName,
				LastName:  &req.LastName,
			}
			return uow.GetUserDetailRepository().CreateUserDetail(ctx, userDetailEntity)
		}

		return nil
	})

	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}

	otp := fmt.Sprintf("%06d", rand.Intn(1000000))
	err = cache.Set(ctx, cache.WithPrefix(KEY_CACHE_OTP_SIGNUP, req.Email), otp, 10*time.Minute)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}

	fmt.Printf("OTP for %s: %s\n", req.Email, otp)

	res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	return res, nil
}

func (s *userService) VerifySignUp(ctx context.Context, req userdto.VerifySignUpCommand) (*userdto.VerifySignUpResponse, error) {
	res := &userdto.VerifySignUpResponse{
		StatusMessage: &statusmsg.StatusMessage{},
	}

	cache := s.Cabin.GetInfra().GetCache()
	exist, err := s.Cabin.GetUnitOfWork().GetUserRepository().GetUserByEmail(ctx, req.Email)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	if exist == nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_USER_NOT_EXIST
		return res, nil
	}
	if exist.Active {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_USER_IS_ACTIVE
		return res, nil
	}

	var otp string
	err = cache.Get(ctx, cache.WithPrefix(KEY_CACHE_OTP_SIGNUP, req.Email), &otp)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, nil
	}
	if otp != req.Otp {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_INCORRECT_OTP
		return res, nil
	}

	err = s.Cabin.GetUnitOfWork().ExecTx(ctx, func(uow repository.UnitOfWork) error {
		exist.Active = true
		return uow.GetUserRepository().UpdateUser(ctx, exist)
	})

	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, nil
	}

	res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	return res, nil
}

func (s *userService) RefreshToken(ctx context.Context, req userdto.RefreshTokenCommand) (*userdto.RefreshTokenResponse, error) {
	res := &userdto.RefreshTokenResponse{
		StatusMessage: &statusmsg.StatusMessage{},
	}

	cache := s.Cabin.GetInfra().GetCache()

	claims, err := auth_context.VerifyJWT(req.RefreshToken, s.Cabin.GetInfra().GetConfig().SecretKey+"_REFRESH_TOKEN")
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_REFRESH_TOKEN_INVALID
		return res, err
	}

	exist, err := s.Cabin.GetUnitOfWork().GetUserRepository().GetUserByGuid(ctx, claims.IdAuthUser)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}
	if exist == nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_USER_NOT_EXIST
		return res, nil
	}
	if !exist.Active {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_USER_NOT_ACTIVE
		return res, nil
	}

	token, err := auth_context.GenerateJWT(&auth_context.ClaimModel{
		IdSite:     exist.SiteId,
		IdAuthUser: exist.Guid,
		Roles:      []string{"user"},
		UserName:   exist.Email,
		Email:      exist.Email,
	}, s.Cabin.GetInfra().GetConfig().SecretKey, 1*time.Hour)
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}

	ttlRefreshToken := 30 * 24 * time.Hour
	refreshToken, err := auth_context.GenerateJWT(&auth_context.ClaimModel{
		IdSite:     exist.SiteId,
		IdAuthUser: exist.Guid,
	}, s.Cabin.GetInfra().GetConfig().SecretKey+"_REFRESH_TOKEN", ttlRefreshToken)

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

func (s *userService) Logout(ctx context.Context, req userdto.LogoutCommand) (*userdto.LogoutResponse, error) {
	res := &userdto.LogoutResponse{
		StatusMessage: &statusmsg.StatusMessage{},
	}

	authContext, ok := helpers.AuthContext(ctx)
	if !ok {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, errors.New("cannot get auth context")
	}

	cache := s.Cabin.GetInfra().GetCache()
	err := cache.Del(ctx, cache.WithPrefix(KEY_CACHE_REFRESH_TOKEN, authContext.IdAuthUser))
	if err != nil {
		res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_UNSPECIFIED
		return res, err
	}

	res.StatusMessage.Code = statusmsg.StatusCode_STATUS_CODE_SUCCESS
	return res, nil
}
