package userusecase

import (
	"errors"
)

var (
	ErrUserNotFound        = errors.New("user_not_found")
	ErrUserNotActive       = errors.New("user_not_active")
	ErrPasswordIncorrect   = errors.New("password_is_incorrect")
	ErrUserExist           = errors.New("user_exist")
	ErrOTPIncorrect        = errors.New("otp_is_incorrect")
	ErrRefreshTokenInvalid = errors.New("refresh_token_is_invalid")
)
